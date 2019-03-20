/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2019/2/16
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package vision

import (
	"github.com/KylinHe/aliensboot-server/module/room/game/agar/protocol"
)

const (
	visiableWidth = 1024

)

func newBlock(mgr *VisionMgr, x int32, y int32) *Block {
	return &Block{
		X:float64(x),
		Y:float64(y),
		mgr:mgr,
		Objs:make(map[*VisionObject]struct{}),
		ObServers:make(map[*VisionUser]struct{}),
	}
}


//显示区域块
type Block struct {
	X float64
	Y float64
	mgr *VisionMgr
	Objs map[*VisionObject]struct{}     //对象
	ObServers map[*VisionUser]struct{} //观察者

}

//新增视野内的对象
func (self *Block) Add(obj *VisionObject) { //ball
	self.Objs[obj] = struct{}{}
	//log.Debug("add block || block x:", self.X, "y:", self.Y)
	for observer := range self.ObServers {
		isRealUser := observer.UserProxy.IsRealUser()
		if isRealUser && observer.UserProxy.HasPlayer() {
			//if observer.ViewObjs[obj] == nil {
			if observer.ViewObjs[obj] == nil {
				observer.ViewObjs[obj] = &ViewObj{EnterSee:true, Ref:1}
				if observer.BeginSee == nil {
					observer.BeginSee =[]*protocol.BallInfo{}
				}
				observer.BeginSee = append(observer.BeginSee, obj.Proxy.PackOnBeginSee())
			} else {
				t := observer.ViewObjs[obj]
				t.Ref = t.Ref +1
			}
		} else if !isRealUser {
			if observer.ViewObjs[obj] == nil {
				observer.ViewObjs[obj] = &ViewObj{ Ref:1}
			} else {
				t := observer.ViewObjs[obj]
				t.Ref = t.Ref + 1
			}
		}
	}
}

//删除视野内的对象
func (self *Block) Remove(obj *VisionObject) { //ball
	delete(self.Objs, obj)
	for observer := range self.ObServers {
		isRealUser := observer.UserProxy.IsRealUser()
		if (isRealUser && observer.UserProxy.HasPlayer())|| !isRealUser {
			if observer.ViewObjs[obj] != nil {
				t := observer.ViewObjs[obj]
				t.Ref = t.Ref - 1
			}
		}
	}
}

func (self *Block) AddObserver(objUser *VisionUser) { //battleUser
	self.ObServers[objUser] = struct {}{}
	//log.Debug("add observer observer length:", len(self.ObServers))
	//log.Debug("self Objs length:", len(self.Objs))
	for obj := range self.Objs {
		isRealUser := objUser.UserProxy.IsRealUser()
		if isRealUser && objUser.UserProxy.HasPlayer() {
			if objUser.ViewObjs[obj] != nil {
				t := objUser.ViewObjs[obj]
				t.Ref = t.Ref + 1
			} else {
				objUser.ViewObjs[obj] = &ViewObj{EnterSee: true, Ref:1}
				if objUser.UserProxy.IsRealUser() && objUser.UserProxy.HasPlayer() {
					if objUser.BeginSee == nil {
						objUser.BeginSee = []*protocol.BallInfo{}
					}
					//observer.Proxy.PackOnBeginSee(objUser.BeginSee)
					objUser.BeginSee = append(objUser.BeginSee, obj.Proxy.PackOnBeginSee())
					//log.Debugf("111")
				}
			}
		} else if !isRealUser {
			if objUser.ViewObjs[obj] != nil {
				t := objUser.ViewObjs[obj]
				t.Ref = t.Ref + 1
			} else {
				objUser.ViewObjs[obj] = &ViewObj{Ref: 1}
			}
		}
	}

}

func (self *Block) RemoveObserver(objUser *VisionUser) {

	//self.ObServers[obj] = nil
	_, ok := self.ObServers[objUser]
	if ok {
		delete(self.ObServers, objUser)
		isRealUser := objUser.UserProxy.IsRealUser()
		if (isRealUser && objUser.UserProxy.HasPlayer()) || !isRealUser {
			for obj := range self.Objs {
				if objUser.ViewObjs[obj] != nil {
					t :=objUser.ViewObjs[obj]
					t.Ref = t.Ref - 1
				}
			}
		}
	}
}