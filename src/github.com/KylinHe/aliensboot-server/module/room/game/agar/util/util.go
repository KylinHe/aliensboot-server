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

import "math"


type MapBorder struct {
	bottomLeft Position
	topRight Position
}

type Position struct {
	X int32
	Y int32
}

func (this Position) Equals(p Position) bool {
	return this.X == p.X && this.Y == p.Y
}

func (this Position) Distance(p Position) float32 {
	var xx = this.X - p.X
	var yy = this.Y - p.Y
	return float32(math.Sqrt(float64(xx * xx + yy * yy)))
}
func (this Position) DistancePow2(p Position) float32 {
	var xx = this.X - p.X
	var yy = this.Y - p.Y
	return xx * xx + yy * yy
}

func NewVector2D(x float64, y float64) *Vector2D {
	return &Vector2D{x:x, y:y}
}

type Vector2D struct {
	x float64
	y float64
}

func (self *Vector2D) Dev(p float64) *Vector2D {
	return NewVector2D(self.x / p, self.y / p)
}

//--向量去模
func (self *Vector2D) Mag() float64 {
	return math.Sqrt(self.x * self.x + self.y * self.y)
}

//--标准化向量
func (self *Vector2D) Normalize() *Vector2D {
	length := math.Sqrt(self.x * self.x + self.y * self.y)
	return self.Dev(length)
}

//--向量点乘
func (self *Vector2D) DotProduct(other Vector2D) *Vector2D {
	return NewVector2D(self.x * other.x, self.y * other.y)
}

//
func (self Vector2D) Copy() *Vector2D {
	return NewVector2D(self.x, self.y)
}

