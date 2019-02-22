/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2019/2/15
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package agar

import (
	"github.com/KylinHe/aliensboot-server/data"
	"github.com/KylinHe/aliensboot-server/module/room/game/agar/util"
)

func NewBall(id int32, owner *BattleUser, otype util.ObjType, pos util.Position, score int32, color int32) *Ball {
	o := &Ball{}
	o.owner = owner
	o.pos = pos
	o.score = score
	o.r = data.AtarGame.Score2R(score)
	o.color = color
	o.id = id
	o.otype = otype
	o.otherVelocitys = []interface{}{}
	o.reqDirection = 0
	o.v = util.NewVector2D(0,0)
	owner.balls[id] = o
	owner.ballCount = owner.ballCount + 1
	o.clientR = o.r
	o.clientPos = pos
	o.bornTick = owner.battle.tickCount

	//owner.battle.colMgr:Enter(o)
	o.collPos = pos
	o.collR = o.r
	return o
}

type Ball struct {
	id int32  //球的唯一id
	owner *BattleUser  //球的拥有者
	score int32 //分数
	r  float64
	v  *util.Vector2D
	pos   util.Position //位置
	color int32  //颜色id
	otype util.ObjType //球类型

	otherVelocitys []interface{}

	reqDirection int32 //请求间隔时间

	clientR float64
	clientPos util.Position
	bornTick int32

	collPos util.Position
	collR float64
}

func (self *Ball) OnDead() {
	if self.otype == util.ObjThorn {
		//self.owner.battle
	}
}
