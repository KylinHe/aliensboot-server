package util

import (
	"math"
)

type MinHeap struct {
	m_size int32
	m_data map[int32]IMinData
}

//type MinData struct {
//	//index int32
//	//timeout int32
//	Proxy IMinData
//}

type IMinData interface {
	GetTimeout() 	int32
	GetIndex()		int32
	SetIndex(index int32)
}

func NewMinHeap() *MinHeap {
	o := &MinHeap{}
	o.m_size = 0
	o.m_data = make(map[int32]IMinData)
	return o
}

func (self *MinHeap) Up(index int32){
	parent_idx := self.Parent(index)
	for parent_idx > 0 {
		if self.m_data[index].GetIndex() < self.m_data[parent_idx].GetIndex() {
			self.swap(index, parent_idx)
			index = parent_idx
			parent_idx = self.Parent(index)
		} else {
			break
		}
	}
}

func (self *MinHeap) Down(index int32){
	l := self.Left(index)
	r := self.Right(index)
	min := index

	if l <= self.m_size && self.m_data[l].GetTimeout() < self.m_data[index].GetTimeout() {
		min = l
	}

	if r <= self.m_size && self.m_data[r].GetTimeout() < self.m_data[min].GetTimeout() {
		min = r
	}

	if min != index {
		self.swap(index, min)
		self.Down(min)
	}
}

func (self *MinHeap) Parent(index int32) int32{
	parent, _ := math.Modf(float64(index)/2)
	return int32(parent)
}

func (self *MinHeap) Left(index int32) int32{
	return 2*index
}

func (self *MinHeap) Right(index int32) int32{
	return 2*index + 1
}

//func (self *MinHeap) Change(o) {
//	index := o.index
//	if index == 0 {
//		return
//	}
//	self.Down(index)
//	if index == o.index {
//		self.Up(index)
//	}
//}

func (self *MinHeap) Insert(o IMinData) {

	if o.GetIndex() != 0 {
		return
	}
	self.m_size = self.m_size + 1
	//self.m_data = append(self.m_data, o)
	self.m_data[self.m_size] = o
	o.SetIndex(self.m_size)
	self.Up(self.m_size)
}

func (self *MinHeap) Min() int32 {
	if self.m_size == 0 {
		return 0
	}
	return self.m_data[1].GetTimeout()
}

func (self *MinHeap) PopMin() IMinData {
	o := self.m_data[1]
	self.swap(1, self.m_size)
	self.m_data[self.m_size] = nil
	self.m_size = self.m_size - 1
	self.Down(1)
	o.SetIndex(0)
	return o
}

func (self *MinHeap) Size() int32 {
	return self.m_size
}

func (self *MinHeap) swap(idx1, idx2 int32) {
	tmp := self.m_data[idx1]
	self.m_data[idx1] = self.m_data[idx2]
	self.m_data[idx2] = tmp

	self.m_data[idx1].SetIndex(idx1)
	self.m_data[idx2].SetIndex(idx2)
}

func (self *MinHeap) Clear() {
	for self.m_size > 0 {
		self.PopMin()
	}
	self.m_size = 0
}