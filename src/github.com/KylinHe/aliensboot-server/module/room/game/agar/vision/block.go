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



const (
	visiableWidth = 1024

)

func newBlock(mgr *Manager, x int32, y int32) *Block {
	return &Block{
		X:x,
		Y:y,
		mgr:mgr,
	}
}


//显示区域块
type Block struct {
	X int32
	Y int32
	mgr *Manager
	Objs map[*VObject]struct{}      //对象
	ObServers map[*VObject]struct{} //观察者

}

//新增视野内的对象
func (self *Block) Add(obj *VObject) {
	self.Objs[obj] = struct{}{}

	for observer, _ := range self.ObServers {
		isRealUser := observer.IsRealUser()
		if isRealUser {
			_, ok := obj.viewObjs[obj]

			if ok {
				observer.enterSee = true
				observer.ref = 1

			} else {

			}
		}
	}
}

//删除视野内的对象
func (self *Block) Remove(obj VObject) {
	self.Objs[obj] = struct{}{}
	for k, _ := range self.ObServers {
		isRealUser := k.IsRealUser()
		if isRealUser {

		}
	}
}



func (self *Block) AddObserver(obj IVObject) {
	self.ObServers[obj] = struct{}{}
	for k, v := range self.Objs {

	}

}
