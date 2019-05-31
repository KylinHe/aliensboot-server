package agar

import (
	util2 "github.com/KylinHe/aliensboot-core/common/util"
	"github.com/KylinHe/aliensboot-server/module/room/game/agar/collision"
	protocol2 "github.com/KylinHe/aliensboot-server/module/room/game/agar/protocol"
	"github.com/KylinHe/aliensboot-server/module/room/game/agar/util"
	"math"
)

type Star struct {
	id		int32
	pos		util.Position
	otype   util.ObjType
	r		float64
	index   int32
	mgr		*StarMgr
	timeout int32

	collision *collision.CollideObject
}


func NewStar(id int32, mgr *StarMgr) *Star {
	o := &Star{}
	o.id = id
	o.mgr = mgr
	s := util.Stars[id]
	o.pos = util.Position{X:float64(s.X), Y:float64(s.Y)}
	o.otype = util.ObjStar
	o.r = 1
	o.index = 0
	o.collision = &collision.CollideObject{
		Proxy:o,
	}
	mgr.room.colMgr.Enter(o.collision)
	return o
}

//碰撞监听
func (self *Star) OnOverLap(other *collision.CollideObject) {
	//log.Debug("Star On Over Lap")
}

func (self *Star) GetPosition() util.Position {
	return self.pos
}

func (self *Star) GetType() util.ObjType {
	return self.otype
}

func (self *Star) GetR() float64 {
	return self.r
}

func (self *Star) Relive () {
	offset1 := int(math.Floor(float64(self.id - 1) / 32))
	offset2 := uint((self.id - 1) % 32)
	// 关闭星星标记
	self.mgr.starBits[offset1 - 1] = self.mgr.starBits[offset1 - 1] | (1 << offset2)
	self.mgr.room.colMgr.Enter(self.collision)
}

type StarMgr struct {
	room *Game
	starBits []uint32
	stars []*Star
	minheap *util.MinHeap
	deads []int32
}

func NewStarMgr(room *Game) *StarMgr{
	o := &StarMgr{}
	cc := len(util.Stars) / 32
	o.starBits = make([]uint32,cc)

	o.room = room
	o.minheap = util.NewMinHeap()
	//for i = 1,cc do
	//table.insert(o.starBits,0xffffffff)
	for i := 0; i < cc; i++ {
		o.starBits[i] = 0xffffffff
	}
	starCount := cc * 32
	//o.stars = make([]*Star, starCount + 1)
	//o.stars[0] = nil
	for i := 0; i < starCount; i++ {
		o.stars = append(o.stars, NewStar(int32(i),o))
	}
	return o
}

func (self *StarMgr) GetStarBits() []uint32 {
	return self.starBits
}

func (self *StarMgr) Update()  {
	nowTick := self.room.tickCount
	reliveStars := []int32{}
	for self.minheap.Size() > 0 {
		if self.minheap.Min() > nowTick {
			break
		} else {
			star := self.minheap.PopMin().(*Star)
			star.Relive()
			reliveStars = append(reliveStars, star.id)
		}
	}

	if len(reliveStars) > 0 {
		msg := &protocol2.StarReliveResp{Cmd:"StarRelive", Stars:reliveStars, Timestamp:self.room.tickCount}
		self.room.Broadcast(msg)
		//self.room.BroadcastOtherPlayer(-1, constant.RoleAll, msg)
	}

	self.NotifyDead()
}

func (self *StarMgr) OnStarDead(star *Star) {
	//从碰撞管理器中移除
	self.room.colMgr.Leave(star.collision)
	offset1 := int32(math.Floor((float64(star.id) - 1)/ 32))
	offset2 := uint((star.id - 1) % 32)
	//关闭星星标记
	//log.Debug("offset1:", offset1)
	self.starBits[offset1 - 1] = self.starBits[offset1 - 1] ^(1 << offset2)
	//复活时间
	star.timeout = self.room.tickCount + int32(util2.RandRange(5000, 8000))
	//log.Debug("relive time:",star.timeout)
	self.minheap.Insert(star)

	self.deads = append(self.deads, star.id)
}

func (self *StarMgr) NotifyDead() {
	if self.deads != nil && len(self.deads) > 0 {
		msg := protocol2.StarDeadResp{Cmd:"StarDead", Stars:self.deads, Timestamp:self.room.tickCount}
		self.room.Broadcast(msg)
		//self.room.BroadcastOtherPlayer(-1, constant.RoleAll, &protocol.Response{})
		self.deads = nil
	}
}


func (self *Star) GetTimeout() int32 {
	return self.timeout
}

func (self *Star) GetIndex() int32 {
	return self.index
}

func (self *Star) SetIndex(index int32){
	self.index = index
}