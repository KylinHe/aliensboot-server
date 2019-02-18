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
)

func newManager() *Manager {
	result := &Manager{}

	xCount := data.AtarGame.MapWidth / data.AtarGame.BlockWidth
	hasRemain := data.AtarGame.MapWidth % data.AtarGame.BlockWidth != 0
	if hasRemain {
		xCount += 1
	}

	result.blocks = make([][]*Block, xCount + 1)
	//区域用的正方形、宽高一致
	for y:=int32(1); y<=xCount; y++ {
		result.blocks[y] = make([]*Block, xCount + 1)
		for x:=int32(1); x<=xCount; x++ {
			result.blocks[y][x] = newBlock(result,x,y)
		}
	}

}

type Manager struct {

	blocks [][]*Block //二位坐标块

}

//获取坐标点所在的块
func (self *Manager) getBLockByPoint(pos util.Position) *Block {
	var x = float64(pos.X)
	var y = float64(pos.Y)
	x = math.Max(0, x)
	y = math.Max(0, y)
	x = math.Min(float64(data.AtarGame.MapWidth-1), x)
	y = math.Min(float64(data.AtarGame.MapWidth-1), y)
	indexX := int(math.Floor(x / float64(data.AtarGame.BlockWidth)))
	indexY := int(math.Floor(y / float64(data.AtarGame.BlockWidth)))
	return self.blocks[indexY+1][indexX+1]
}


func (self *Manager) calBlocks(pos util.Position) *Block {
	var x = float64(pos.X)
	var y = float64(pos.Y)
	x = math.Max(0, x)
	y = math.Max(0, y)
	x = math.Min(float64(data.AtarGame.MapWidth-1), x)
	y = math.Min(float64(data.AtarGame.MapWidth-1), y)
	indexX := int(math.Floor(x / float64(data.AtarGame.BlockWidth)))
	indexY := int(math.Floor(y / float64(data.AtarGame.BlockWidth)))
	return self.blocks[indexY+1][indexX+1]
}



