/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2019/2/15
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package collision

import (
	"github.com/KylinHe/aliensboot-server/data"
	"github.com/KylinHe/aliensboot-server/module/room/game/agar/util"
)


func NewQuadTreeCollision() *QuadTreeCollision {
	return &QuadTreeCollision{QuadTree:NewQuadTree(NewRect(0,0, data.AtarGame.MapWidth, data.AtarGame.MapWidth))}
}


func rect(pos util.Position, r float64) *Rect {
	var bottomLeftX = pos.X - r
	if bottomLeftX < 0 {
		bottomLeftX = 0
	}

	var bottomLeftY = pos.Y - r
	if bottomLeftY < 0 {
		bottomLeftY = 0
	}

	var topRightX = pos.X + r
	if topRightX > data.AtarGame.MapWidth {
		topRightX = data.AtarGame.MapWidth
	}


	var topRightY = pos.Y + r
	if topRightY > data.AtarGame.MapWidth {
		topRightY = data.AtarGame.MapWidth
	}

	return NewRect(bottomLeftX,bottomLeftY,topRightX,topRightY)
}

type QuadTreeCollision struct {
	*QuadTree
}

func (self *QuadTreeCollision) Enter(obj *CollideObject) {
	if obj.Collision != nil {
		return
	}
	//log.Debug("Enter Collision || obj type:", obj.Proxy.GetType())
	obj.Rect = rect(obj.Proxy.GetPosition(), obj.Proxy.GetR())
	obj.Collision = self
	self.QuadTree.Insert(obj)
}

func (self *QuadTreeCollision) Leave(obj *CollideObject) {
	if obj.Collision != nil {
		obj.Tree.Remove(obj)
		obj.Collision = nil
	}
}

func (self *QuadTreeCollision) Update(obj *CollideObject) {
	if obj.Collision != nil {
		obj.Rect = rect(obj.Proxy.GetPosition(), obj.Proxy.GetR())
		self.QuadTree.Update(obj)
	}
}

func checkCollision(obj *CollideObject, other *CollideObject) {
	if other.Proxy != nil {
		totalR := obj.Proxy.GetR() + other.Proxy.GetR()
		distancePow2 := obj.Proxy.GetPosition().DistancePow2(other.Proxy.GetPosition())
		if distancePow2 < totalR * totalR {
			//log.Debugf("obj  ||type:%v, pos:%v, r:%v",obj.Proxy.GetType(), obj.Proxy.GetPosition(), obj.Proxy.GetR())
			//log.Debugf("other||type:%v, pos:%v, r:%v",other.Proxy.GetType(), other.Proxy.GetPosition(), other.Proxy.GetR())
			obj.Proxy.OnOverLap(other)
		}
	}
}

func (self *QuadTreeCollision) CheckCollision(obj *CollideObject) {
	if obj.Collision != nil {
		self.QuadTree.RectCall(obj.Rect, func (other *CollideObject) {
			if other != obj {
				checkCollision(obj, other)
			}
		})
	}
}

