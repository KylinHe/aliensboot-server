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
	util2 "github.com/KylinHe/aliensboot-core/common/util"
	"github.com/KylinHe/aliensboot-server/data"
	protocol2 "github.com/KylinHe/aliensboot-server/module/room/game/agar/protocol"
	"github.com/KylinHe/aliensboot-server/module/room/game/agar/util"
	"github.com/KylinHe/aliensboot-server/module/room/game/agar/vision"
	"github.com/KylinHe/aliensboot-server/protocol"
	"math"
	"sort"
)

/**
 * 玩家战斗对象
 */
type BattleUser struct {
	//*entity.Player
	Player *Player

	userID int64

	color int32 //皮肤id

	BallCount int32 //当前球的数量

	Balls map[int32]*Ball //拥有的球

	Battle *Game

	ViewPort *util.MapBorder

	//viewObjs map[*vision.VisionObject]*ViewObj

	VisionCenter util.Position

	//visionObj *vision.VisionObject
	visionUser *vision.VisionUser

	stop bool

	reqDirection float64


	//beiginsee map[int32]*Ball
}

func NewBattleUser(player *Player, userID int64, battle *Game) *BattleUser {
	o := &BattleUser{}
	o.userID = userID
	o.color = int32(util2.RandRange(0, len(data.AtarGame.Colors)))
	o.Balls = make(map[int32]*Ball)
	o.BallCount = 0
	if player != nil {
		o.Player = player
		player.battleUser = o
	}
	o.Battle = battle
	o.visionUser = &vision.VisionUser{
		UserProxy:o,
		Vision:battle.visionMgr,
		//Collision: battle.colMgr,
		//: o,
	}
	//battleUser.visionObj = &vision.VisionObject{
	//	Collision:room
	//	Proxy:battleUser,
	//}

	//o.collisionObj = &collisionObj.CollideObject{
	//	Collision:self.battle.colMgr,
	//	OType:thorn.otype,
	//	Position:thorn.pos,
	//	R:thorn.r,
	//	Obj:thorn,
	//}
	//o.userID = userID
	o.stop = true
	o.ViewPort = &util.MapBorder{}
	return o
}

func (self *BattleUser)Relive() {
	if self.userID == 0 {
		return
	}
	r := math.Ceil(data.AtarGame.Score2R(data.AtarGame.InitScore))
	pos := util.Position{}
	mapWidth := self.Battle.mapBorder.TopRight.X - self.Battle.mapBorder.BottomLeft.X
	mapHeight := self.Battle.mapBorder.TopRight.Y - self.Battle.mapBorder.BottomLeft.Y
	pos.X = float64(util2.RandInt32Scop(int32(r), int32(mapWidth -r)))
	pos.Y = float64(util2.RandInt32Scop(int32(r), int32(mapHeight - r)))

	newBall := NewBall(self.Battle.GetBallID(), self, util.ObjBall, pos, data.AtarGame.InitScore, self.color)
	//if newBall != nil {
	//	newBall.visionObj = &vision.VisionObject{
	//		Vision:self.Battle.visionMgr,
	//		Proxy:newBall,
	//		Pos:newBall.pos,
	//		R:newBall.r,
	//	}
	//	newBall.collisionObj = &collision.CollideObject{
	//		Collision:self.Battle.colMgr,
	//		Proxy:newBall,
	//		OType:newBall.otype,
	//		Position:newBall.pos,
	//		R:newBall.r,
	//		Obj:newBall,
	//	}
	//}
	if newBall != nil {
		self.Battle.visionMgr.Enter(newBall.visionObj)
	}
}

func (self *BattleUser) UpdateGatherTogether() {
	if self.userID == 0 {
		return
	}
	if self.stop {
		if self.BallCount > 1 {
			cx := 0.0
			cy := 0.0
			for _, v := range self.Balls {
				cx = cx + v.pos.X
				cy = cy + v.pos.Y
			}
			cx = cx / float64(self.BallCount)
			cy = cy / float64(self.BallCount)
			for _, v := range self.Balls {
				v.GatherTogether(util.Position{X:cx,Y:cy})
			}
		} else {
			for _, v := range self.Balls {
				if v.moveVelocity != nil && v.moveVelocity.V.Mag() > 0 {
					v.Stop()
				}
			}
		}
	}
}

func (self *BattleUser) Update(elapse float64) {
	if self.userID != 0 && self.BallCount == 0 {
		self.Relive()
	}

	if !self.stop {
		self.UpdateBallMovement()
	} else {
		self.UpdateGatherTogether()
	}
	for _, v := range self.Balls {
		v.Update(elapse)
		if self.userID != 0 {

			//log.Debugf("userID:%v || viewPort:%v || visionCenter:%v || pos:%v || reqDirection:%v || V:%v", self.userID, self.ViewPort, self.VisionCenter, v.pos,self.reqDirection, v.v)
			self.Battle.colMgr.CheckCollision(v.collisionObj)
		}
	}
	if self.userID >= 1000 {
		//目前只有真实玩家使用视野
		self.Battle.visionMgr.UpdateUserVision(self.visionUser)
	}

	for _, v := range self.Balls {
		self.visionUser.Vision.UpdateVisionObj(v.visionObj)
	}
}

func (self *BattleUser) RefreshBallsUpdateInfo() {
	for _, v := range self.Balls {
		if v.otype == util.ObjBall {
			if v.r != v.clientR || !v.pos.Equals(v.clientPos) {
				//log.Debug("client r", v.clientR, "r", v.r, "client pos", v.clientPos, "pos", v.pos)
				v.visionObj.BallUpdateInfo = &util.UpdateBallInfo{
					Id:v.id,
					R:v.r,
					Pos:util.Position{X:v.pos.X, Y:v.pos.Y},
					V: v.PredictV(),
				}
				//if v.predictV == nil {
				//	v.visionObj.BallUpdateInfo.V = v.PredictV()
				//}
				v.clientR = v.r
				v.clientPos = util.Position{X:v.pos.X, Y:v.pos.Y}
			} else {
				v.visionObj.BallUpdateInfo = nil
			}
		}
	}
}

func (self *BattleUser) NotifyBalls2Client(elapse float64) {
	if self.Player != nil {
		ballUpdate := []*util.UpdateBallInfo{}
		endSee := []int32{}
		for k, v := range self.visionUser.ViewObjs {
			if v.Ref <= 0 {
				delete(self.visionUser.ViewObjs, k)
				//self.visionUser.ViewObjs[k] = nil
				endSee = append(endSee, k.Proxy.GetBallID())
			} else if v.EnterSee {
				v.EnterSee = false
			} else if elapse != 0 {
				if k.BallUpdateInfo != nil {
					ballUpdate = append(ballUpdate, k.BallUpdateInfo)
				}
			}
		}
		if self.visionUser.BeginSee != nil && len(self.visionUser.BeginSee) > 0 {
			//log.Debug("response=======================BeginSee")
			resp := &protocol2.BeginSeeResp{Cmd:"BeginSee", Timestamp:self.Battle.tickCount, Balls:self.visionUser.BeginSee}
			self.Battle.Send2Client(resp, self.userID - 1000)
			self.visionUser.BeginSee = []*protocol2.BallInfo{}
		}

		if len(ballUpdate) > 0 {
			//log.Debug("response=======================BallUpdate")
			resp := &protocol2.BallUpdateResp{Cmd:"BallUpdate", Timestamp:self.Battle.tickCount, Balls:ballUpdate}
			self.Battle.Send2Client(resp, self.userID - 1000)
		}

		if len(endSee) > 0 {
			//log.Debug("response=======================EndSee")
			resp := &protocol2.EndSeeResp{Cmd:"EndSee", Timestamp:self.Battle.tickCount, Balls:endSee}
			self.Battle.Send2Client(resp, self.userID - 1000)
		}
	}

}


func (self *BattleUser) Move(msg *Request) {
	self.stop = false
	if self.BallCount == 1 {
		for _, v := range self.Balls {
			//log.Debugf("%v",v)
			v.Move(msg.Dir) //(msg.dir)
		}
	}
	self.reqDirection = msg.Dir// msg.dir
}

func (self *BattleUser) Stop(msg *protocol.Request) {
	self.reqDirection = 0
	self.stop = true
	for _, v := range self.Balls {
		v.Stop()
	}
}

func (self *BattleUser) OnBallDead(ball *Ball) {
	if ball.owner == self {
		delete(self.Balls, ball.id)
		//self.Balls[ball.id] = nil
		self.BallCount = self.BallCount - 1
	}
}

//吐孢子
func (self *BattleUser) Spit() {
	for _, v := range self.Balls {
		v.Spit()
	}
}

//分裂
func (self *BattleUser) Split() {
	balls := []*Ball{}
	for _,v := range self.Balls {
		balls = append(balls, v)
	}
	sort.Sort(BallsSlice(balls))
	//table.sort(balls,function (a,b)
	//return a.r > b.r
	//end)
	for _, v := range balls {
		v.Split()
	}
	self.UpdateBallMovement()

}

type BallsSlice []*Ball

func (a BallsSlice) Len() int {    // 重写 Len() 方法
	return len(a)
}
func (a BallsSlice) Swap(i, j int){     // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a BallsSlice) Less(i, j int) bool {    // 重写 Less() 方法， 从大到小排序
	return a[j].r < a[i].r
}

func (self *BattleUser) UpdateBallMovement() {
	if self.userID == 0 || self.BallCount < 1 || self.reqDirection == 0 {
		return
	} else {
		//先计算小球的几何重心
		cx := 0.0
		cy := 0.0
		for _, v := range self.Balls {
			cx = cx + v.pos.X
			cy = cy + v.pos.Y
		}
		centralPos := util.Position{X:cx/float64(self.BallCount), Y:cy/float64(self.BallCount)}
		maxDis := 0.0
		for _, v := range self.Balls {
			dis := v.pos.Distance(centralPos) + v.r
			if dis > maxDis {
				maxDis = dis
			}
		}
		forwardDis := maxDis + 300

		//vDir := util.NewVector2D()
		forwardPoint := util.MoveTo(centralPos, self.reqDirection, forwardDis, util.Position{}, util.Position{})
		for _, v := range self.Balls {
			vv := util.NewVector2D(forwardPoint.X- v.pos.X, forwardPoint.Y - v.pos.Y)
			v.Move(vv.GetDirAngle())
		}
	}
}

func (self *BattleUser)IsRealUser() bool {//是否真实玩家
	return self.userID > 1000
}

func (self *BattleUser) HasPlayer() bool {//有玩家
	return self.Player != nil
}

func (self *BattleUser) GetBallCount() int32 {
	return self.BallCount
}

func (self *BattleUser) GetBallInfo() []*vision.BallInfo {
	result := []*vision.BallInfo{}
	for _, ball := range self.Balls {
		result = append(result, &vision.BallInfo{R:ball.r, Pos:ball.pos, Score:ball.score})
	}
	return result
}

func (self *BattleUser) GetVisionCenter() util.Position {
	return self.VisionCenter
}

func (self *BattleUser) GetViewPort() *util.MapBorder {
	return self.ViewPort
}

func (self *BattleUser) SetVisionCenter(position util.Position) {
	self.VisionCenter.X = position.X
	self.VisionCenter.Y = position.Y
}

func (self *BattleUser) SetViewPort(bottomLeft util.Position, topRight util.Position){
	self.ViewPort.TopRight = topRight
	self.ViewPort.BottomLeft = bottomLeft
}
//func in_range(top_right util.Position, bottom_left util.Position, x float64, y float64) bool {
//	if x >= bottom_left.X && y >= bottom_left.Y && x <= top_right.X && y <= top_right.Y {
//		return true
//	}
//	return false
//}


