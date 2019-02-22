/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/11/22
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package core

import (
	"github.com/KylinHe/aliensboot-server/constant"
	"github.com/KylinHe/aliensboot-server/protocol"
	"github.com/KylinHe/aliensboot-core/exception"
)

//座位编号从1开始
type Seats []*Seat

func NewSeats(seatNum int) Seats {
	result := make(Seats, seatNum)
	for i := 0; i < seatNum; i++ {
		result[i] = &Seat{id: int32(i) + 1}
	}
	return result
}

//新增玩家
func (seats Seats) Add(player *Player, seatID int32) *Seat {
	if seats.Exists(player.GetId()) {
		exception.GameException(protocol.Code_playerAlreadySeat)
	}
	if seatID == constant.AnySeat {
		seat := seats.allocEmptySeat()
		if seat == nil {
			exception.GameException(protocol.Code_roomMaxSeat)
		}
		seat.SetPlayer(player)
		return seat
	} else {
		seat := seats.EnsureSeat(seatID)
		seat.SetPlayer(player)
		return seat
	}
}

func (seats Seats) UpdateLock(seatID int32, lock bool) *Seat {
	seat := seats.EnsureSeat(seatID)
	seat.lock = lock
	return seat
}

func (seats Seats) EnsureSeat(seatID int32) *Seat {
	seatSeq := int(seatID - 1)
	if seatSeq > len(seats) {
		exception.GameException(protocol.Code_invalidSeat)
	}
	return seats[seatSeq]
}

func (seats Seats) allocEmptySeat() *Seat {
	for _, seat := range seats {
		if seat.IsFree() {
			return seat
		}
	}
	return nil
}

func (seats Seats) Get(playerID int64) *Player {
	for _, seat := range seats {
		if seat.CheckPlayer(playerID) {
			return seat.GetPlayer()
		}
	}
	return nil
}

func (seats Seats) Delete(playerID int64) *Player {
	for _, seat := range seats {
		if seat.CheckPlayer(playerID) {
			return seat.RemovePlayer()
		}
	}
	return nil
}

func (seats Seats) Exists(playerID int64) bool {
	return seats.Get(playerID) != nil
}

//func (seats Seats) IsFull() bool {
//	for _, seat := range seats {
//		if seat.player == nil {
//			return false
//		}
//	}
//	return true
//}

func (seats Seats) GetMaxSeat() int32 {
	return int32(len(seats))
}

func (seats Seats) Clean() {
	for _, seat := range seats {
		seat.RemovePlayer()
	}
}

func (seats Seats) Foreach(callback func(player *Player)) {
	for _, seat := range seats {
		if seat.GetPlayer() != nil {
			callback(seat.GetPlayer())
		}
	}
}

func (seats Seats) ForeachSeat(callback func(seat *Seat)) {
	for _, seat := range seats {
		callback(seat)
	}
}


func (seats Seats) IsAllReady() bool {
	for _, seat := range seats {
		player := seat.GetPlayer()
		if player == nil || !player.IsReady() {
			return false
		}
	}
	return true
}
