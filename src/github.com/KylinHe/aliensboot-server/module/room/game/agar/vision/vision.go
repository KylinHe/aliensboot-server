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
	"github.com/KylinHe/aliensboot-server/module/room/game/agar/collision"
	"github.com/KylinHe/aliensboot-server/module/room/game/agar/protocol"
	"github.com/KylinHe/aliensboot-server/module/room/game/agar/util"
)

func NewVision() IVision {
	return newManager()
}

//visionObject
type VisionObject struct {
	//VisionMgr *VisionMgr

	Vision IVision		//vision mgr

	Collision collision.ICollision //collision mgr

	Proxy IVisionObject //代理实现对象

	//viewObjs map[*VisionObject]*ViewObj //视野内对象    // ball -- {EnterSee, Ref}

	visionBlocks *VisionBlocks

	BallUpdateInfo *util.UpdateBallInfo
}

type VisionUser struct {
	UserProxy IVisionUser                //battle user 代理
	Vision    IVision                    //vision mgr
	ViewObjs  map[*VisionObject]*ViewObj //视野内对象    // ball -- {EnterSee, Ref}
	BeginSee  []*protocol.BallInfo
	Balls     []*BallInfo
}

type BallInfo struct {
	R   float64
	Pos util.Position
	Score int32
}


type ViewObj struct {
	EnterSee bool
	Ref      int
}

//type TT struct {
//	UserID int64
//	Id int32
//	R float64
//	Pos util.Position
//	Color int32
//	Veloctitys []*util.Velocity
//}


type IVisionObject interface {
	//HasPlayer() bool //有玩家
	//
	//IsRealUser() bool //是否真实玩家
	GetPos() util.Position

	PackOnBeginSee() *protocol.BallInfo //

	GetR() float64

	GetBallID() int32
}

type IVisionUser interface {
	HasPlayer() bool //有玩家
	IsRealUser() bool //是否真实玩家
	GetBallCount() int32
	GetBallInfo() []*BallInfo
	GetVisionCenter() util.Position
	GetViewPort() *util.MapBorder

	SetVisionCenter(position util.Position)
	SetViewPort(bottomLeft util.Position, topRight util.Position)
}

type IVision interface {
	Enter(obj *VisionObject)

	Leave(obj *VisionObject)

	UpdateUserVision(obj *VisionUser)

	UpdateVisionObj(obj *VisionObject)
}

