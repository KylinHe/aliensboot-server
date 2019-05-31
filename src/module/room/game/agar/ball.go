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
	"github.com/KylinHe/aliensboot-core/log"
	"github.com/KylinHe/aliensboot-server/data"
	"github.com/KylinHe/aliensboot-server/module/room/game/agar/collision"
	"github.com/KylinHe/aliensboot-server/module/room/game/agar/protocol"
	"github.com/KylinHe/aliensboot-server/module/room/game/agar/util"
	"github.com/KylinHe/aliensboot-server/module/room/game/agar/vision"
	"math"
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
	o.otherVelocitys = make(map[*util.Velocity]struct{})
	o.reqDirection = 0
	o.v = util.NewVector2D(0,0)
	owner.Balls[id] = o
	owner.BallCount = owner.BallCount + 1
	o.clientR = o.r
	o.clientPos = pos
	o.bornTick = owner.Battle.tickCount
	o.collisionObj = &collision.CollideObject{
		//Collision:self.owner.Battle.colMgr,
		Proxy:o,
	}
	o.visionObj = &vision.VisionObject{
		//Vision:owner.Battle.visionMgr, //交给vision初始化
		Collision:owner.Battle.colMgr,
		Proxy:o,
	}
	owner.Battle.colMgr.Enter(o.collisionObj)

	o.collPos = pos
	o.collR = o.r
	return o
}

type Ball struct {
	id           int32  //球的唯一id
	owner        *BattleUser  //球的拥有者
	score        int32 //分数
	r            float64
	v            *util.Vector2D
	pos          util.Position //位置
	color        int32  //颜色id
	otype        util.ObjType //球类型
	splitTimeout float64

	otherVelocitys map[*util.Velocity]struct{}
	moveVelocity *util.Velocity

	reqDirection float64 //请求间隔时间

	clientR float64
	clientPos util.Position
	bornTick int32

	//ballUpdateInfo *util.UpdateBallInfo
	predictV *util.Vector2D

	collisionObj *collision.CollideObject
	visionObj    *vision.VisionObject

	collPos util.Position
	collR float64

	needThorn bool
}

func (this *Ball) GetPos() util.Position {
	return this.pos
}

func (self *Ball) OnDead() {
	if self.otype == util.ObjThorn {
		self.owner.Battle.thornMgr.OnThornDead()
	}
	self.owner.Battle.colMgr.Leave(self.collisionObj)
	self.owner.Battle.visionMgr.Leave(self.visionObj)
	self.owner.OnBallDead(self)
}

func (self *Ball)fixBorder() {
	mapBorder := self.owner.Battle.mapBorder
	bottomLeft := mapBorder.BottomLeft
	topRight := mapBorder.TopRight
	R := self.r * math.Sin(math.Pi/4)
	self.pos.X = math.Max(R+bottomLeft.X, self.pos.X)
	self.pos.X = math.Min(topRight.X-R, self.pos.X)
	self.pos.Y = math.Max(R + bottomLeft.Y, self.pos.Y)
	self.pos.Y = math.Min(topRight.Y - R, self.pos.Y)
}

func (self *Ball) UpdatePosition(averageV *util.Vector2D, elapse float64) {
	elapse = elapse/1000
	self.pos.X = self.pos.X + averageV.GetX() * elapse
	self.pos.Y = self.pos.Y + averageV.GetY() * elapse
	self.fixBorder()
}

//计算一个预测速度
func (self *Ball) PredictV() *util.Vector2D {
	//计算一个预测速度
	var predictVelocitys []*util.Velocity
	if self.moveVelocity != nil {
		predictVelocitys = append(predictVelocitys, self.moveVelocity.Copy())
	}
	for k := range self.otherVelocitys {
		predictVelocitys = append(predictVelocitys, k.Copy())
	}

	predictV := util.NewVector2D(0, 0)
	for _, v := range predictVelocitys {
		predictV = util.Add(predictV, v.Update(float64(self.owner.Battle.tickInterval)))
	}
	predictV = util.Div(predictV, 3)
	return predictV
}

func (self *Ball) Update(elapse float64) {
	if float64(self.owner.Battle.tickCount) > self.splitTimeout {
		self.splitTimeout = 0
	}

	self.v = util.NewVector2D(0,0)

	if self.moveVelocity != nil {
		self.v = self.moveVelocity.Update(elapse)
	}

	for k := range self.otherVelocitys {
		self.v = util.Add(self.v, k.Update(elapse))
		if k.Duration <= 0 {
			//self.otherVelocitys[k] = nil

			//self.otherVelocitys = append(self.otherVelocitys[:k], self.otherVelocitys[k+1:]...)
			delete(self.otherVelocitys, k)
		}
	}

	if self.v.Mag() <= 0 {
		self.moveVelocity = nil
		return
	}

	//更新位置
	self.UpdatePosition(self.v, elapse)
	if self.otype != util.ObjThorn {
		//如果球的半径和坐标自上次更新碰撞之后都没变更过，就不需要再更新碰撞
		if self.collR != self.r || ! self.collPos.Equals(self.pos) {
			self.collR = self.r
			self.collPos = util.Position{X:self.pos.X, Y:self.pos.Y}
			self.owner.Battle.colMgr.Update(self.collisionObj)
		}
	}
	self.ProcessThorn()
}

func (self *Ball) Move(direction float64) {
	speed := data.AtarGame.SpeedByR(self.r, 0)
	self.reqDirection, _ = math.Modf(direction) //60
	maxVelocity := util.TransFormV(self.reqDirection, speed)
	if self.moveVelocity != nil {
		self.moveVelocity = util.NewVelocity(self.moveVelocity.V, maxVelocity, 200, 0 )
	} else {
		self.moveVelocity = util.NewVelocity(util.TransFormV(0, 0), maxVelocity, 200, 0)
	}
}

func (self *Ball) GatherTogether(centerPos util.Position) {
	vv := util.NewVector2D(centerPos.X - self.pos.X, centerPos.Y - self.pos.Y)
	speed := data.AtarGame.SpeedByR(self.GetR(), 0.0 ) * data.AtarGame.CentripetalSpeedCoef
	self.reqDirection= vv.GetDirAngle()
	velocity := util.TransFormV(self.reqDirection, speed)
	self.moveVelocity = util.NewVelocity(velocity, nil, 0 ,0)
}

func (self *Ball) Stop() {
	if self.moveVelocity != nil {
		self.moveVelocity = util.NewVelocity(self.moveVelocity.V, util.NewVector2D(0,0), 200,200)
	}
}

func calSplitTimeout(score int32) float64 {
	return math.Floor(math.Sqrt(float64(score) * 4))*1000
}

func (self *Ball) EatStar(star *Star) {
	self.owner.Battle.starMgr.OnStarDead(star)
	self.score = self.score + data.AtarGame.StarScore
	self.r = data.AtarGame.Score2R(self.score)
	if !self.owner.stop && self.moveVelocity != nil {
		speed := data.AtarGame.SpeedByR(self.GetR(), 0)
		//将传入的角度和速度标量转换成一个速度向量
		maxVelocity := util.TransFormV(self.reqDirection, speed)
		self.moveVelocity = util.NewVelocity(self.moveVelocity.V, maxVelocity, 200, 0)
	}
}

func (self *Ball) EatSpore(other *Ball) {
	other.OnDead()
	self.score = self.score + other.score
	self.r = data.AtarGame.Score2R(self.score)
	if !self.owner.stop && self.moveVelocity != nil {
		speed := data.AtarGame.SpeedByR(self.GetR(), 0)
		maxVelocity := util.TransFormV(self.reqDirection, speed)
		self.moveVelocity = util.NewVelocity(self.moveVelocity.V, maxVelocity, 200,0)
	}
}

func (self *Ball) ProcessThorn() {
	if !self.needThorn {
		return
	}
	self.needThorn = false
	eatFactor := data.AtarGame.EatFactor(float64(self.score))
	n1 := math.Min(float64(data.AtarGame.MaxUserBallCount - self.owner.BallCount), float64(data.AtarGame.MaxThornBallCount))
	x := float64(self.score) / n1
	var n, S2 float64
	if x < float64(data.AtarGame.InitScore) * eatFactor {
		n = float64(self.score) / (float64(data.AtarGame.InitScore) *eatFactor)
		S2 = float64(data.AtarGame.InitScore) * eatFactor
	} else if x <= float64(data.AtarGame.InitScore) * (eatFactor + 2) {
		n = n1
		S2 = float64(self.score)/n
	} else {
		n = n1
		S2 = float64(data.AtarGame.InitScore) * (2 + eatFactor)
	}
	splitCount := int32(n - 1)
	if splitCount == 0 || self.owner.BallCount + splitCount > data.AtarGame.MaxUserBallCount {
		self.r = data.AtarGame.Score2R(self.score)
		if !self.owner.stop && self.moveVelocity != nil {
			speed := data.AtarGame.SpeedByR(self.r, 0)
			//将传入的角度和速度标量转换成一个速度向量
			maxVelocity := util.TransFormV(self.reqDirection, speed)
			self.moveVelocity = util.NewVelocity(self.moveVelocity.V, maxVelocity, 200, 0)
		}
		return
	}

	delta := int32(math.Floor(360/float64(splitCount)))
	L := 8 * data.AtarGame.ScreenSizeFactor
	v0 := data.AtarGame.SpeedByR(data.AtarGame.Score2R(int32(S2)), 0) * data.AtarGame.SpitV0Factor
	spitDuration := int32(math.Floor((2*L/v0) * 1000))

	_scoreRemain := float64(self.score) - S2 * float64(splitCount)

	var i int32 = 0
	for splitCount > 0 {
		self.spit(self.owner, util.ObjBall, int32(S2), int32(_scoreRemain), float64(i), v0, spitDuration)
		i = i + delta
		splitCount --
	}
	self.splitTimeout = float64(self.owner.Battle.tickCount) + calSplitTimeout(self.score)
}

func (self *Ball) EatThorn(thorn *Ball){
	thorn.OnDead()
	self.score = self.score + thorn.score
	if self.owner.BallCount < data.AtarGame.MaxUserBallCount {
		self.needThorn = true
		//不能再这里直接调用ProcessThorn,会导致collision中迭代出错
		//（这个函数是在collision的迭代中调进来的,thorn:OnDead会从collision中移除元素，如果同时调用ProcessThorn向collision中添加元素
		//就会导致迭代器出错）
	} else {
		self.r = data.AtarGame.Score2R(self.score)
		if !self.owner.stop && self.moveVelocity != nil {
			speed := data.AtarGame.SpeedByR(self.r, 0)
			maxVelocity := util.TransFormV(self.reqDirection, speed)
			self.moveVelocity = util.NewVelocity(self.moveVelocity.V, maxVelocity, 200, 0)
		}
	}
}

func (self *Ball) EatBall (other *Ball) {
	other.OnDead()
	self.score = self.score + other.score
	self.r = data.AtarGame.Score2R(self.score)
	if self.owner == other.owner {
		self.splitTimeout = float64(self.owner.Battle.tickCount) + calSplitTimeout(self.score)
	}
	if !self.owner.stop && self.moveVelocity != nil {
		speed := data.AtarGame.SpeedByR(self.r, 0)
		//将传入的角度和速度标量转换成一个速度向量
		maxVelocity := util.TransFormV(self.reqDirection, speed)
		self.moveVelocity = util.NewVelocity(self.moveVelocity.V, maxVelocity, 200, 0)
	}
}

func canEat(b1, b2 *Ball) bool {
	eatFactor := data.AtarGame.EatFactor(float64(b1.score))
	if float64(b1.score)/ float64(b2.score) >= eatFactor {
		return true
	} else {
		return false
	}
}

type CellCollision struct {
	totalR	float64
	dx 		float64
	dy 		float64
	squared float64
}

func checkCellCollision(ball1, ball2 *Ball) *CellCollision{
	totalR := ball1.r + ball2.r
	dx := ball2.pos.X - ball1.pos.X
	dy := ball2.pos.Y - ball1.pos.Y
	squared := dx * dx + dy * dy
	if squared > totalR * totalR {
		return nil
	} else {
		return &CellCollision{
			totalR:totalR,
			dx:dx,
			dy:dy,
			squared:squared,
		}
	}
}

func (self *Ball) OnSelfBallOverLap (other *Ball) {
	manifold := checkCellCollision(self, other)
	if manifold != nil {
		ball1 := self
		ball2 := other
		d := math.Sqrt(manifold.squared)
		if d <= 0 {
			return
		}

		invd := 1 / d
		nx := math.Floor(manifold.dx) * invd
		ny := math.Floor(manifold.dy) * invd
		penetration := (manifold.totalR - d) * 0.75
		if penetration <= 0 {
			return
		}

		px := penetration * nx
		py := penetration * ny

		totalMass := ball1.score + ball2.score
		if totalMass <= 0 {
			return
		}

		invTotalMass := 1 / totalMass
		impulse1 := ball2.score * invTotalMass
		impulse2 := ball1.score * invTotalMass

		ball1.pos.X = ball1.pos.X - (px * float64(impulse1))
		ball1.pos.Y = ball1.pos.Y - (py * float64(impulse1))
		ball2.pos.X = ball2.pos.X + (px * float64(impulse2))
		ball2.pos.Y = ball2.pos.Y + (py * float64(impulse2))

		ball1.fixBorder()
		ball2.fixBorder()
	}
}

func addCollisionElasticity(ball1, ball2 *Ball) {
	dir1To2 := util.NewVector2D(ball2.pos.X-ball1.pos.X, ball2.pos.Y-ball1.pos.Y).GetDirAngle()
	dir2To1, _ := math.Modf(dir1To2 + 180) //360
	manifold := checkCellCollision(ball1, ball2)
	if manifold != nil {
		d := math.Sqrt(manifold.squared)
		if d <= 0 {
			return
		}
		// invd := 1/ d
		//nx := math.Floor(manifold.dx) * invd
		//ny := math.Floor(manifold.dy) * invd
		penetration := (manifold.totalR - d) * 0.75
		if penetration <= 0 {
			return
		}
		//px := penetration * nx
		//py := penetration * ny

		totalMass := ball1.score + ball2.score
		if totalMass <= 0 {
			return
		}

		invTotalMass := 1 / totalMass
		impulse1 := ball2.score * invTotalMass
		impulse2 := ball1.score * invTotalMass

		v1 := util.TransFormV(dir2To1, ball1.v.Mag()*0.8*float64(impulse2))
		velocity1 := util.NewVelocity(v1, util.TransFormV(0, 0), 200, 200)
		//ball1.otherVelocitys = append(ball1.otherVelocitys, velocity1)
		ball1.otherVelocitys[velocity1] = struct{}{}

		v2 := util.TransFormV(dir1To2, ball1.v.Mag()*0.8*float64(impulse1))
		velocity2 := util.NewVelocity(v2, util.TransFormV(0, 0), 200, 200)
		//ball2.otherVelocitys = append(ball2.otherVelocitys, velocity2)
		ball2.otherVelocitys[velocity2] = struct{}{}

		ball1.bornTick = 0
		ball2.bornTick = 0
	}
}

func (self *Ball)spit(owner *BattleUser, newtype util.ObjType, spitScore int32,spitterScore int32, dir float64, v0 float64, duration int32) *Ball{
	spitR := data.AtarGame.Score2R(spitScore)
	leftBottom := util.Position{X:spitR, Y:spitR}
	rightTop := util.Position{X:data.AtarGame.MapWidth - spitR, Y:data.AtarGame.MapWidth - spitR}
	spiterR := data.AtarGame.Score2R(spitterScore)
	//log.Debug("spiterR:", spiterR)
	//log.Debug("spitR:", spitR)
	bornPoint := util.MoveTo(self.pos, dir, spiterR + spitR, leftBottom, rightTop)
	//log.Debug("pos:",self.pos, "bornPoint:",bornPoint)
	var color int32
	if newtype == util.ObjBall {
		color = self.color
	} else {
		color = util2.RandInt32Scop(0, int32(len(data.AtarGame.Colors)))
	}
	self.score = spitterScore
	self.r = spiterR

	newBall := NewBall(self.owner.Battle.GetBallID(), owner, newtype, bornPoint, spitScore, color)

	if newtype == util.ObjBall {
		newBall.splitTimeout = float64(self.owner.Battle.tickCount) + calSplitTimeout(newBall.score)
	}

	//增加弹射运动量
	velocity := util.NewVelocity(util.TransFormV(dir, v0), util.TransFormV(0, 0 ), float64(duration), float64(duration))
	newBall.otherVelocitys[velocity] = struct{}{}
	if !self.owner.stop {
		if newtype == util.ObjBall {
			speed := data.AtarGame.SpeedByR(spitR, 0)
			//将传入的角度和速度标量转换成一个速度向量
			maxVelocity := util.TransFormV(self.reqDirection, speed)
			newBall.moveVelocity = util.NewVelocity(maxVelocity, nil, 0 ,0)
		}
		//自己的积分减少，速度改变了
		speed := data.AtarGame.SpeedByR(self.r, 0)
		//将传入的角度和速度标量转换成一个速度向量
		maxVelocity := util.TransFormV(self.reqDirection, speed)
		self.moveVelocity = util.NewVelocity(self.moveVelocity.V, maxVelocity, 200, 0)
	}
	self.owner.Battle.visionMgr.Enter(newBall.visionObj)

	return newBall
}

func (self *Ball) Spit() {
	eatFactor := data.AtarGame.EatFactor(float64(self.score))
	if float64(self.score) >= float64(data.AtarGame.Sp0) * (1 + eatFactor) {
		log.Debug("spit")
		spitR := data.AtarGame.Score2R(data.AtarGame.Sp0)
		L := 9 * data.AtarGame.ScreenSizeFactor
		v0 := data.AtarGame.SpeedByR(spitR, 0) * data.AtarGame.SpitV0Factor
		spitDuration := math.Floor((2*L/v0)*1000)
		self.spit(self.owner.Battle.dummyUser, util.ObjSpore, data.AtarGame.Sp0, self.score - data.AtarGame.Sp0, self.reqDirection, v0, int32(spitDuration))
	}
}

func (self *Ball) splitAble() bool {
	eatFactor := data.AtarGame.EatFactor(float64(self.score))
	if float64(self.score) < float64(data.AtarGame.Sp0) * eatFactor * 2 {
		return false
	}
	return true
}

func (self *Ball)Split() {
	if self.owner.BallCount >= data.AtarGame.MaxUserBallCount {
		return
	}
	if !self.splitAble() {
		return
	}
	newR := data.AtarGame.Score2R(self.score/2)
	L := newR + 5.5 * data.AtarGame.ScreenSizeFactor
	v0 := math.Floor(2 * L * 1000 / float64(data.AtarGame.SplitDuration))
	self.spit(self.owner, util.ObjBall, self.score/2, self.score/2, self.reqDirection, v0, data.AtarGame.SplitDuration)
	self.splitTimeout = float64(self.owner.Battle.tickCount) + calSplitTimeout(self.score)
}

func (self *Ball) PackOnBeginSee() *protocol.BallInfo {
	//log.Debug("PackOnBeginSee")
	tt := &protocol.BallInfo{}
	tt.UserID = self.owner.userID
	tt.ID = self.id
	tt.R = self.r
	tt.Pos = util.Position{X:self.pos.X, Y:self.pos.Y}
	tt.Color = self.color
	var velocitys []*util.Velocity
	if self.moveVelocity != nil {
		velocitys = append(velocitys, self.moveVelocity.Pack())
	}
	if self.otherVelocitys != nil {
		for k := range self.otherVelocitys {
			velocitys = append(velocitys, k.Pack())
		}
	}

	if len(velocitys) > 0 {
		for _, velocity := range velocitys {
			tt.Velocitys = append(tt.Velocitys, &protocol.Velocity{
				Duration:velocity.Duration,
				AccRemain:velocity.AccRemain,
				V:&protocol.Vector2D{X:velocity.V.GetX(), Y:velocity.V.GetY()},
				TargetV:&protocol.Vector2D{X:velocity.TargetV.GetX(), Y:velocity.TargetV.GetY()},
				//A:&protocol.Vector2D{X:velocity.A.GetX(), Y:velocity.A.GetY()},
			})
		}
	}
	//t = append(t, tt)
	//log.Debug("222")
	return tt

}

func (self *Ball) GetPosition() util.Position {
	return self.pos
}

func (self *Ball) GetType() util.ObjType {
	return self.otype
}

func (self *Ball) GetR() float64 {
	return self.r
}

func (self *Ball) GetBallID() int32{
	return self.id
}

func (self *Ball) OnOverLap(other *collision.CollideObject)  {
	//log.Debug("Ball OnOverLap || other type:", other.Proxy.GetType())
	if self.otype == util.ObjSpore {
		return
	}
	if other.Proxy.GetType() == util.ObjStar {
		self.EatStar(other.Proxy.(*Star))
	} else if other.Proxy.GetType() == util.ObjSpore {
		distance := self.pos.Distance(other.Proxy.GetPosition())
		if distance <= self.r && canEat(self, other.Proxy.(*Ball)) {
			self.EatSpore(other.Proxy.(*Ball))
		}
	} else if other.Proxy.GetType() == util.ObjBall {
		if  other.Proxy.(*Ball).owner == self.owner {
			if  other.Proxy.(*Ball).splitTimeout > 0  || self.splitTimeout > 0{
				if self.bornTick > 0  && self.owner.Battle.tickCount < self.bornTick + data.AtarGame.SplitDuration {
					//增加碰撞弹射运动量
					addCollisionElasticity(self, other.Proxy.(*Ball))
				} else {
					self.OnSelfBallOverLap(other.Proxy.(*Ball))
				}
			} else {
				distance := self.pos.Distance(other.Proxy.GetPosition())
				if distance <= self.r {
					self.EatBall(other.Proxy.(*Ball))
				}
			}
		} else {
			distance := self.pos.Distance(other.Proxy.GetPosition())
			if distance <= self.r && canEat(self, other.Proxy.(*Ball)) {
				self.EatBall(other.Proxy.(*Ball))
			}
		}
	} else if other.Proxy.GetType() == util.ObjThorn {
		distance := self.pos.Distance(other.Proxy.GetPosition())
		if distance <= self.r &&  canEat(self, other.Proxy.(*Ball)) {
			self.EatThorn(other.Proxy.(*Ball))
		}
	}
}
