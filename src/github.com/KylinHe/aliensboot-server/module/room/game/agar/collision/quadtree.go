/*******************************************************************************
 * Co.Yright (c) 2015, 2017 aliens idea.Xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2019/2/15
 * Contributors:
 *     aliens idea.Xiamen) Corporation - initial API and implementation
 *     jialin.he <.Ylinh@gmail.com>
 *******************************************************************************/
package collision

import (
	"github.com/KylinHe/aliensboot-core/log"
	"github.com/KylinHe/aliensboot-server/data"
	"math"
)



func NewQuadTree(rect *Rect) *QuadTree {
	return new(0, rect, 1)
}

func new(index int32, rect *Rect, level int32) *QuadTree {
	return &QuadTree{
		rect:rect,
		objs:make(map[*CollideObject]struct{}),
		objCount:0,
		index:index,
		level:level,
	}
}

type QuadTree struct {
	rect *Rect
	nodes map[int32]*QuadTree //四元树子树

	objs map[*CollideObject]struct{} //子对象
	objCount int32                   //包含的对象数量

	level int32
	index int32
}

//根据范围获取数据
func (self *QuadTree) getSubTree(rect *Rect) *QuadTree {
	if self.nodes == nil {
		return nil
	}
	for _, subTree := range self.nodes {
		if subTree.rect.Include(rect) {
			return subTree
		}
	}
	return nil
}

func (self *QuadTree) Split() {
	var  subWidth = float32(math.Ceil(float64(self.rect.Width()) / 2))
	var  subHeight = float32(math.Ceil(float64(self.rect.Height()) / 2))

	var bottomLeft = self.rect.bottomLeft
	var topRight = self.rect.topRight

	self.nodes = make(map[int32]*QuadTree, 4)

	self.nodes[1] = new(1, NewRect(bottomLeft.X, bottomLeft.Y + subHeight, bottomLeft.X + subWidth, topRight.Y),self.level + 1)
	self.nodes[2] = new(2, NewRect(bottomLeft.X + subWidth, bottomLeft.Y + subHeight, topRight.X, topRight.Y),self.level + 1)
	self.nodes[3] = new(3, NewRect(bottomLeft.X, bottomLeft.Y, bottomLeft.X + subWidth, bottomLeft.Y + subHeight),self.level + 1)
	self.nodes[4] = new(4, NewRect(bottomLeft.X + subWidth, bottomLeft.Y, topRight.X, bottomLeft.Y + subHeight),self.level + 1)
}

func (self *QuadTree) Insert(obj *CollideObject) bool {
	if obj.Tree != nil {
		log.Error("obj.tree != nil")
		return false
	}
	//不在范围内不能吃
	if !self.rect.Include(obj.Rect) {
		return false
	}
	var subTree = self.getSubTree(obj.Rect)
	if subTree != nil {
		subTree.Insert(obj)
	}

	//球的子对象达到阈值 需要分裂
	if self.objCount + 1 > data.AtarGame.MaxObjs && self.level < data.AtarGame.MaxLevel && self.nodes == nil {
		self.Split()
	}

	obj.Tree = self
	self.objs[obj] = struct{}{}
	self.objCount += 1
	return true
}


//--获取与rect相交的空间内的所有对象
//func (self *QuadTree) Retrieve(rect *Rect, objs []*QuadTree) {
//	if !self.rect.Intersect(rect) {
//		return
//	}
//	if self.nodes != nil {
//		for _, node := range self.nodes {
//			node.Retrieve(rect, objs)
//		}
//	}
//
//	for _, obj := range self.objs {
//		objs = append(objs, obj)
//	}
//}

//--对每个与rect相交的空间内的对象执行func
func (self *QuadTree) RectCall(rect *Rect, callback func(tree *CollideObject)) {
	if !self.rect.Intersect(rect) {
		return
	}
	if self.nodes != nil {
		for _, node := range self.nodes {
			node.RectCall(rect, callback)
		}
	}
	for obj, _ := range self.objs {
		callback(obj)
	}
}

func (self *QuadTree) Remove(obj *CollideObject) {
	if obj.Tree == self {
		_, ok := self.objs[obj]
		if ok {
			delete(self.objs, obj)
			obj.Tree = nil
			self.objCount -= 1
		}
	} else {
		log.Error("obj.tree != self")
	}
}


func (self *QuadTree) Update(obj *CollideObject) {
	if self.level != 1 {
		log.Error("update should call in level == 1")
		return
	}

	var tree = obj.Tree
	if tree != nil {
		//需要执行更新的条件
		//1 当前子树不能完全容纳obj.rect
		//2 当前子树有任意一个子节点能完全容纳obj.rect
		for {
			if !tree.rect.Include(obj.Rect) {
				break
			}

			var nodes = tree.nodes
			if nodes != nil && (nodes[1].rect.Include(obj.Rect) ||
				nodes[2].rect.Include(obj.Rect)  || nodes[3].rect.Include(obj.Rect)  || nodes[4].rect.Include(obj.Rect)) {
					break
			}
			return
		}

		tree.Remove(obj)
		self.Insert(obj)
	}

}



