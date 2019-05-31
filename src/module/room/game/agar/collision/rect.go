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
	"github.com/KylinHe/aliensboot-core/log"
	"github.com/KylinHe/aliensboot-server/module/room/game/agar/util"
)


func NewRect(bottomLeftX float64,bottomLeftY float64,topRightX float64,topRightY float64) *Rect {
	return &Rect{
		bottomLeft: util.Position{X:bottomLeftX, Y:bottomLeftY},
		topRight: util.Position{X:topRightX, Y:topRightY},
	}
}


func inRange(topRight util.Position, bottomLeft util.Position, x float64, y float64) bool {
	return  x >= bottomLeft.X && y >= bottomLeft.Y && x <= topRight.X && y <= topRight.Y
}

//
type Rect struct {
	topRight 	util.Position
	bottomLeft 	util.Position
}

func (self *Rect) Height() float64 {
	return self.topRight.Y - self.bottomLeft.Y
}

func (self *Rect) Width() float64 {
	if self == nil {
		log.Error(self)
	}
	return self.topRight.X - self.bottomLeft.X
}

//是否交叉
func (self *Rect) Intersect(other *Rect) bool {
	var oTopLeft = util.Position{X: other.topRight.X - other.Width(), Y:other.topRight.Y}
	var oBottomRight = util.Position{X:other.bottomLeft.X + other.Width(), Y:other.bottomLeft.Y}

	if inRange(self.topRight,self.bottomLeft,oTopLeft.X,oTopLeft.Y) {
		return true
	}

	if inRange(self.topRight,self.bottomLeft,oBottomRight.X,oBottomRight.Y) {
		return true
	}

	if inRange(self.topRight,self.bottomLeft,other.topRight.X,other.topRight.Y) {
		return true
	}

	if inRange(self.topRight,self.bottomLeft,other.bottomLeft.X,other.bottomLeft.Y) {
		return true
	}

	return false
}

//--返回与other是否完全被包含
func (self *Rect) Include(other *Rect) bool {
	var oTopLeft = util.Position{X: other.topRight.X - other.Width(), Y:other.topRight.Y}
	var oBottomRight = util.Position{X:other.bottomLeft.X + other.Width(), Y:other.bottomLeft.Y}

	if !inRange(self.topRight,self.bottomLeft,oTopLeft.X,oTopLeft.Y) {
		return false
	}

	if !inRange(self.topRight,self.bottomLeft,oBottomRight.X,oBottomRight.Y) {
		return false
	}

	if !inRange(self.topRight,self.bottomLeft,other.topRight.X,other.topRight.Y) {
		return false
	}

	if !inRange(self.topRight,self.bottomLeft,other.bottomLeft.X,other.bottomLeft.Y) {
		return false
	}

	return true
}



