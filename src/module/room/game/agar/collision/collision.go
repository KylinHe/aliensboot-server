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

import "github.com/KylinHe/aliensboot-server/module/room/game/agar/util"

func NewCollision() ICollision {
	return NewQuadTreeCollision()
}

type CollideObject struct {

	//OType util.ObjType

	Collision ICollision  //mgr

	Tree *QuadTree

	Rect *Rect

	//Position util.Position

	//R float64

	//Obj interface{}

	Proxy ICollideObject

}

type ICollideObject interface {

	GetPosition() util.Position

	GetType() util.ObjType

	GetR() float64
	//碰撞监听
	OnOverLap(other *CollideObject)

}

type ICollision interface {

	Enter(obj *CollideObject)

	Leave(obj *CollideObject)

	Update(obj *CollideObject)

	CheckCollision(obj *CollideObject)

}

