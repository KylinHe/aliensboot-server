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
	"github.com/KylinHe/aliensboot-server/data"
	"github.com/KylinHe/aliensboot-server/module/room/game/agar/util"
	"math"
	"reflect"
)

var blockWidth = data.AtarGame.MapWidth/50

const (
	visibleSizeWidth = 1024
	visibleSizeHeight = 768
)


type VisionMgr struct {

	Blocks [][]*Block //二位坐标块

}

func newManager() *VisionMgr {
	result := &VisionMgr{}

	xCount := data.AtarGame.MapWidth/ float64(data.AtarGame.BlockWidth)
	hasRemain := int32(data.AtarGame.MapWidth) % data.AtarGame.BlockWidth != 0
	if hasRemain {
		xCount += 1
	}
	xCount = math.Floor(xCount)
	result.Blocks = make([][]*Block, int32(xCount)+ 1)
	//区域用的正方形、宽高一致
	for y:=int32(1); y<=int32(xCount); y++ {
		result.Blocks[y] = make([]*Block, int32(xCount) + 1)
		for x:=int32(1); x<=int32(xCount); x++ {
			result.Blocks[y][x] = newBlock(result,x,y)
		}
	}
	return result
}

//获取坐标点所在的块
func (self *VisionMgr) GetBLockByPoint(pos util.Position) *Block {
	var x = float64(pos.X)
	var y = float64(pos.Y)
	x = math.Max(0, x)
	y = math.Max(0, y)
	x = math.Min(float64(data.AtarGame.MapWidth-1), x)
	y = math.Min(float64(data.AtarGame.MapWidth-1), y)
	indexX := int(math.Floor(x / float64(data.AtarGame.BlockWidth)))
	indexY := int(math.Floor(y / float64(data.AtarGame.BlockWidth)))
	return self.Blocks[indexY+1][indexX+1]
}


type VisionBlocks struct {
	block_info	*BlockInfo
	blocks		[]*Block
}

type BlockInfo struct {
	top_left	 util.Position
	top_right	 util.Position
	bottom_left  util.Position
	bottom_right util.Position
}

func (self *VisionMgr) calBlocks(obj *VisionObject) *VisionBlocks {
	blocks := &VisionBlocks{}
	bottom_left := util.Position{
		X:math.Max(0, obj.Proxy.GetPos().X - obj.Proxy.GetR()),
		Y:math.Max(0, obj.Proxy.GetPos().Y - obj.Proxy.GetR()),
	}
	top_right := util.Position{
		X:math.Min(data.AtarGame.MapWidth - 1, obj.Proxy.GetPos().X + obj.Proxy.GetR()),
		Y:math.Min(data.AtarGame.MapWidth - 1, obj.Proxy.GetPos().Y + obj.Proxy.GetR()),
	}
	block_bottom_left := self.GetBLockByPoint(bottom_left)
	block_top_right := self.GetBLockByPoint(top_right)

	blocks.block_info = &BlockInfo{}
	blocks.block_info.bottom_left = util.Position{X: float64(block_bottom_left.X), Y:float64(block_bottom_left.Y)}
	blocks.block_info.top_right = util.Position{X: float64(block_top_right.X), Y: float64(block_top_right.Y)}

	blocks.blocks = []*Block{}
	for y := block_bottom_left.Y; y <= block_top_right.Y; y++ {
		for x := block_bottom_left.X; x <= block_top_right.X; x++ {
			blocks.blocks = append(blocks.blocks, self.Blocks[int(y)][int(x)])
		}
	}

	return blocks
}

//对象进入视野系统
func (self *VisionMgr) Enter(obj *VisionObject) { //ball
	//log.Debugf("enter vision, ball id:%v",obj.Proxy.GetBallID())
	obj.visionBlocks = self.calBlocks(obj)
	for _, v := range obj.visionBlocks.blocks {
		v.Add(obj)
	}
	obj.Vision = self
}

//对象离开视野系统
func (self *VisionMgr) Leave(obj *VisionObject) { //ball
	if obj.Vision != nil {
		for _, v := range obj.visionBlocks.blocks {
			v.Remove(obj)
		}
		obj.visionBlocks = nil
		obj.Vision = nil
	}
}

func (self *VisionMgr) UpdateVisionObj(obj *VisionObject) { //

	if obj.Collision != nil {
		//计算出新管理块
		blocks := self.calBlocks(obj)

		for _, v := range obj.visionBlocks.blocks {
			//从离开的管理块移除对象(在老单元,不在新单元范围内的块)
			if !in_range(blocks.block_info.top_right, blocks.block_info.bottom_left, v.X, v.Y) {
				v.Remove(obj)
				//log.Debug("111 remove obj x:", , "y:", y)
			}
		}

		for _, v := range blocks.blocks {
			//向新加入的管理块加入对象（在新单元,不在老管理单元范围内的块）
			if !in_range(obj.visionBlocks.block_info.top_right,obj.visionBlocks.block_info.bottom_left, v.X, v.Y) {
				v.Add(obj)
			}
		}
		obj.visionBlocks = blocks
	}
}

func in_range(top_right util.Position, bottom_left util.Position, x float64, y float64) bool {
	if x >= bottom_left.X && y >= bottom_left.Y && x <= top_right.X && y <= top_right.Y {
		return true
	}
	return false
}

//更新玩家视野
func (self *VisionMgr) UpdateUserVision(obj *VisionUser) {
	//user := obj//.(*agar.BattleUser)
	if obj.UserProxy.IsRealUser() && !obj.UserProxy.HasPlayer(){
		return
	}
	//首先计算玩家视野中心点
	if obj.UserProxy.GetBallCount() > 0 {
		cx := 0.0
		cy := 0.0
		for _, v := range obj.UserProxy.GetBallInfo(){
			cx = cx + v.Pos.X
			cy = cy + v.Pos.Y
		}
		cx = cx / float64(obj.UserProxy.GetBallCount())
		cy = cy / float64(obj.UserProxy.GetBallCount())
		//obj.VisionCenter = util.Position{X:cx, Y:cy}
		obj.UserProxy.SetVisionCenter(util.Position{X: cx, Y:cy})
	}

	//log.Debug("view port:",obj.UserProxy.GetViewPort())

	//计算视野范围
	bottomLeft, topRight := self.updateViewPort(obj)
	//obj.ViewPort.BottomLeft = bottomLeft
	//obj.ViewPort.TopRight = topRight
	if reflect.DeepEqual(obj.UserProxy.GetViewPort().BottomLeft, util.Position{}) || reflect.DeepEqual(obj.UserProxy.GetViewPort().TopRight, util.Position{}) {
		obj.UserProxy.SetViewPort(bottomLeft, topRight)
		//log.Debug("111 new view port:",bottomLeft, topRight)

		for y := bottomLeft.Y; y <= topRight.Y; y ++ {
			for x := bottomLeft.X; x <= topRight.X; x ++ {
				//log.Debug("111 add observer x:", x, "y:", y)
				self.Blocks[int(y)][int(x)].AddObserver(obj)
			}
 		}
	} else {

		oldBottomLeft := obj.UserProxy.GetViewPort().BottomLeft
		oldTopRight := obj.UserProxy.GetViewPort().TopRight

		if oldBottomLeft.X != bottomLeft.X || oldBottomLeft.Y != bottomLeft.Y ||
			oldTopRight.X != topRight.X || oldTopRight.Y != topRight.Y {
			obj.UserProxy.SetViewPort(bottomLeft, topRight)
			//log.Debug("222 new view port:",bottomLeft, topRight)

			for y := oldBottomLeft.Y; y <= oldTopRight.Y; y ++ {
				for x := oldBottomLeft.X; x <= oldTopRight.X; x ++ {
					// 在老单元，不在新单元中
					if !in_range(topRight, bottomLeft, x, y) {
						//log.Debug("222 remove observer x:", x, "y:", y)
						self.Blocks[int(y)][int(x)].RemoveObserver(obj)
					}
				}
			}
			for y := bottomLeft.Y; y <= topRight.Y; y++ {
				for x := bottomLeft.X; x <= topRight.X; x++ {
					if !in_range(oldTopRight, oldBottomLeft, x, y) {
						// 在新单元，不在老单元中
						//log.Debug("222 add observer x:", x, "y:", y)
						self.Blocks[int(y)][int(x)].AddObserver(obj)
					}
				}
			}
		}
	}
}

func (self *VisionMgr)updateViewPort(user *VisionUser) (util.Position, util.Position) {
	if  user.UserProxy.GetBallCount() == 0 {
		return user.UserProxy.GetViewPort().BottomLeft, user.UserProxy.GetViewPort().TopRight
	}

	_edgeMaxX := 0.0
	_edgeMaxY := 0.0
	_edgeMinX := 1000000.0
	_edgeMinY := 1000000.0

	for _, v := range user.UserProxy.GetBallInfo() {
		R := math.Floor(data.AtarGame.Score2R(v.Score))
		bottomLeft := util.Position{X:v.Pos.X - R, Y:v.Pos.Y - R}
		topRight := util.Position{X:v.Pos.X + R, Y:v.Pos.Y + R}

		if _edgeMaxX < topRight.X {
			_edgeMaxX = topRight.X
		}

		if _edgeMaxY < topRight.Y {
			_edgeMaxY = topRight.Y
		}

		if _edgeMinX > bottomLeft.X {
			_edgeMinX = bottomLeft.X
		}

		if _edgeMinY > bottomLeft.Y {
			_edgeMinY = bottomLeft.Y
		}
	}

	width := _edgeMaxX - _edgeMinX
	height := _edgeMaxY - _edgeMinY

	para := 30.0
	r := math.Max(width, height)
	r = (r*0.5) / para

	a1 := 8 / math.Sqrt(r)
	a2 := math.Max(a1, 1.5)
	a3 := r * a2
	a4 := math.Max(a3, 10)
	a5 := math.Min(a4, 100)
	scale := a5 * para

	scale = scale / (float64(data.AtarGame.VisibleHeight) / 2)

	//log.Debug("scale:",scale)

	_visionWidth := math.Floor(float64(data.AtarGame.VisibleWidth) * scale + 300)
	_visionHeight := math.Floor(float64(data.AtarGame.VisibleHeight) * scale + 300)

	bottomLeft := util.Position{}
	bottomLeft.X = math.Max(1, user.UserProxy.GetVisionCenter().X - _visionWidth/2)
	bottomLeft.Y = math.Max(1, user.UserProxy.GetVisionCenter().Y - _visionHeight/2)

	block := self.GetBLockByPoint(bottomLeft)
	bottomLeft.X = block.X
	bottomLeft.Y = block.Y

	topRight := util.Position{}
	topRight.X = math.Min(data.AtarGame.MapWidth - 1, user.UserProxy.GetVisionCenter().X + _visionWidth/2)
	topRight.Y = math.Min(data.AtarGame.MapWidth - 1, user.UserProxy.GetVisionCenter().Y + _visionHeight/2)

	block1 := self.GetBLockByPoint(topRight)
	topRight.X = block1.X
	topRight.Y = block1.Y

	return bottomLeft, topRight
}



