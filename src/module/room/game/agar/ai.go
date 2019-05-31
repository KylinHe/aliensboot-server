package agar

import (
	"github.com/KylinHe/aliensboot-core/common/util"
)

type AiMgr struct {
	battle *Game
	robots map[int64]*Robot
}

type Robot struct {
	nextMove int32
	user *BattleUser
}

func NewAiMgr(battle *Game, aiCount int32) *AiMgr {
	o := &AiMgr{}
	o.battle = battle
	o.robots = make(map[int64]*Robot)
	for i := 1; int32(i) <= aiCount ; i++ {
		robot := &Robot{}
		robot.user = NewBattleUser(nil, int64(i), battle)
		//robot.user.Battle = battle
		//robot.user.visionObj = &vision.VisionObject{
		//	Vision:battle.visionMgr,
		//	Proxy:robot.user,
		//}
		robot.user.Relive()
		battle.SetUser(int32(i), robot.user)
		robot.nextMove = battle.GetTickCount()
		o.robots[int64(i)] = robot
	}
	return o
}

func (self *AiMgr) Update() {
	for _, v := range self.robots {
		if self.battle.GetTickCount() > v.nextMove {
			v.user.Move(&Request{Dir:float64(util.RandInt32Scop(0,359))})
			v.nextMove = self.battle.GetTickCount() + util.RandInt32Scop(4000, 6000)
		}
	}
}