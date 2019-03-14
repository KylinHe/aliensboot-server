/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2019/2/15
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package util

import (
	"github.com/KylinHe/aliensboot-core/log"
	"math"
	"reflect"
	"time"
)


type MapBorder struct {
	BottomLeft Position
	TopRight   Position
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (this Position) Equals(p Position) bool {
	return this.X == p.X && this.Y == p.Y
}

//两点距离
func (this Position) Distance(p Position) float64 {
	var xx = this.X - p.X
	var yy = this.Y - p.Y
	return math.Sqrt(xx * xx + yy * yy)
}

//两点距离的平方
func (this Position) DistancePow2(p Position) float64 {
	var xx = this.X - p.X
	var yy = this.Y - p.Y
	return xx * xx + yy * yy
}

func MoveTo(p Position, dir float64, distance float64, leftBottom Position, topRight Position) Position {
	target := Position{X:p.X, Y:p.Y}
	rad := math.Pi/180*dir
	target.X = target.X + distance * math.Cos(rad)
	target.Y = target.Y + distance * math.Sin(rad)

	if !reflect.DeepEqual(leftBottom, Position{}){
		log.Debug("leftBottom is not nil")
		target.X = math.Max(leftBottom.X, target.X)
		target.Y = math.Max(leftBottom.Y, target.Y)
	}

	if !reflect.DeepEqual(topRight, Position{}) {
		log.Debug("topRight is not nil")
		target.X = math.Min(topRight.X, target.X)
		target.Y = math.Min(topRight.Y, target.Y)
	}

	return target
}

func NewVector2D(x float64, y float64) *Vector2D {
	return &Vector2D{X: x, Y:y}
}

type Vector2D struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (self *Vector2D) GetX() float64 {
	return self.X
}

func (self *Vector2D) GetY() float64 {
	return self.Y
}

func (self *Vector2D) Dev(p float64) *Vector2D {
	return NewVector2D(self.X/ p, self.Y/ p)
}

//--向量去模
func (self *Vector2D) Mag() float64 {
	return math.Sqrt(self.X* self.X + self.Y* self.Y)
}

//--标准化向量
func (self *Vector2D) Normalize() *Vector2D {
	length := math.Sqrt(self.X* self.X + self.Y* self.Y)
	return self.Dev(length)
}

//--向量点乘
func (self *Vector2D) DotProduct(other Vector2D) *Vector2D {
	return NewVector2D(self.X* other.X, self.Y* other.Y)
}

func (self *Vector2D) copy() *Vector2D {
	return NewVector2D(self.X, self.Y)
}

func (self *Vector2D) GetDirAngle() float64 {
	return math.Mod(math.Atan2(self.Y, self.X)*180 / math.Pi + 360,360)
}

//--向量相减
func Sub(p1, p2 *Vector2D) *Vector2D {
	return NewVector2D(p1.X- p2.X, p1.Y- p2.Y)
}

//--向量相加
func Add(p1, p2 *Vector2D) *Vector2D {
	return NewVector2D(p1.X+ p2.X, p1.Y+p2.Y)
}

//--向量乘num
func Mul(p1 *Vector2D, p2 float64) *Vector2D {
	return NewVector2D(p1.X* p2, p1.Y* p2)
}

//--向量除num
func Div(p1 *Vector2D, p2 float64) *Vector2D {
	return NewVector2D(p1.X/ p2, p1.Y/ p2)
}



type Velocity struct {
	runTime   float64
	Duration  float64
	AccRemain float64
	A         *Vector2D
	TargetV   *Vector2D
	V         *Vector2D
}

func NewVelocity(v0, v1 *Vector2D, accelerateTime, duration float64) *Velocity {
	result := &Velocity{}
	if v1 == nil {
		v1 = v0
	}
	if duration == 0 {
		result.Duration = math.MaxUint32
	}
	result.Duration = duration
	result.AccRemain = 0
	if v0 != v1 && accelerateTime > 0 {
		//变速运动
		result.V = v0.copy()
		result.A = Div(Sub(v1, v0), accelerateTime / 1000)
		result.TargetV = v1.copy()
		result.AccRemain = accelerateTime
	} else {
		//匀速运动
		result.TargetV = v0.copy()
		result.V = v0.copy()
	}
	if result.TargetV == nil {
		log.Debug("targetV is nil")
	}
	result.Duration = math.Max(result.Duration, result.AccRemain)
	return result
}

func (self *Velocity) Update(elapse float64) *Vector2D {
	if self.Duration == 0 {
		return NewVector2D(0,0)
	}
	self.runTime = self.runTime + elapse
	deltaAcc := math.Min(elapse, self.AccRemain)
	self.AccRemain = self.AccRemain - deltaAcc
	delta := math.Min(elapse, self.Duration)
	self.Duration = self.Duration - delta

	if self.AccRemain > 0 {
		lastV := self.V.copy()
		self.V = Add(self.V, Mul(self.A, deltaAcc/1000))
		return Div(Add(lastV, self.V),2)
	} else {
		backV := self.V.copy()
		self.V = self.TargetV.copy()
		if deltaAcc > 0 {
			return Add(Mul(Div(Add(backV, self.TargetV), 2), deltaAcc/elapse), Mul(self.TargetV, (delta-deltaAcc)/elapse))
		} else {
			return Mul(Div(Add(backV, self.TargetV), 2), delta/elapse)
		}
	}

}

func (self *Velocity) Pack() *Velocity {
	//velocitys := []*Velocity{}
	velocity := &Velocity{}
	velocity.AccRemain = self.AccRemain
	velocity.Duration = self.Duration
	velocity.V = NewVector2D(self.V.X, self.V.Y) //{x: self.TargetV.x, y:self.TargetV.y}
	velocity.TargetV = NewVector2D(self.TargetV.X, self.TargetV.Y)
	//velocitys = append(velocitys, velocity)
	return velocity
}

func (self *Velocity) Copy() *Velocity {
	return NewVelocity(self.V, self.TargetV, self.AccRemain, self.Duration)
}

func TransFormV(direction float64, v float64) *Vector2D {
	direction, _ = math.Modf(direction)
	rad := math.Pi/180.0*direction
	return NewVector2D(math.Cos(rad) * v, math.Sin(rad) * v)
}

func TimeNowSysTick() int32 {
	return int32(time.Now().UnixNano() / 1e6)
}

type UpdateBallInfo struct {
	Id int32	`json:"id"`
	R  float64 	`json:"r"`
	Pos Position `json:"pos"`
	V  *Vector2D `json:"v"`
}