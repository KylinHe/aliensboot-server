package agar

import (
	util2 "github.com/KylinHe/aliensboot-core/common/util"
	"github.com/KylinHe/aliensboot-server/data"
	"github.com/KylinHe/aliensboot-server/module/room/game/agar/util"
	"math"
)

type ThornMgr struct {
	battle *Game
}

func NewThornMgr (battle *Game) *ThornMgr {
	o := &ThornMgr{}
	o.battle = battle
	for i := int32(0); i <= data.AtarGame.ThornCount; i++ {
		score := util2.RandInt32Scop(data.AtarGame.MinThornScore, data.AtarGame.MaxThornScore)
		r := math.Ceil(data.AtarGame.Score2R(score))
		pos := util.Position{
			X: float64(util2.RandInt32Scop(int32(r), int32(data.AtarGame.MapWidth - r))),
			Y: float64(util2.RandInt32Scop(int32(r), int32(data.AtarGame.MapWidth - r))),
		}
		thorn := NewBall(battle.GetBallID(), battle.dummyUser, util.ObjThorn, pos, score, data.AtarGame.ThornColorID)
		//thorn.visionObj = &vision.VisionObject{
		//	Vision:battle.visionMgr,
		//	Proxy:thorn,
		//	Pos:thorn.pos,
		//	R:thorn.r,
		//}
		//thorn.collisionObj = &collision.CollideObject{
		//	Collision:battle.colMgr,
		//	Proxy:thorn,
		//	OType:thorn.otype,
		//	Position:thorn.pos,
		//	R:thorn.r,
		//	Obj:thorn,
		//}
		o.battle.visionMgr.Enter(thorn.visionObj)
	}
	return o
}

func (self *ThornMgr) OnThornDead() {
	score := util2.RandInt32Scop(data.AtarGame.MinThornScore, data.AtarGame.MaxThornScore)
	r := math.Ceil(data.AtarGame.Score2R(score))
	pos := util.Position{
		X:float64(util2.RandInt32Scop(int32(r), int32(data.AtarGame.MapWidth - r))),
		Y:float64(util2.RandInt32Scop(int32(r), int32(data.AtarGame.MapWidth - r))),
	}
	thorn := NewBall(self.battle.GetBallID(), self.battle.dummyUser, util.ObjThorn, pos, score, data.AtarGame.ThornColorID)
	//thorn.visionObj = &vision.VisionObject{
	//	Vision:self.battle.visionMgr,
	//	Proxy:thorn,
	//	Pos:thorn.pos,
	//	R:thorn.r,
	//}
	//thorn.collisionObj = &collision.CollideObject{
	//	Collision:self.battle.colMgr,
	//	Proxy:thorn,
	//	OType:thorn.otype,
	//	Position:thorn.pos,
	//	R:thorn.r,
	//	Obj:thorn,
	//}
	self.battle.visionMgr.Enter(thorn.visionObj)
}
