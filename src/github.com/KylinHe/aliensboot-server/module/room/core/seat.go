/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/11/13
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

type Seat struct {

	id int32 //座位编号

	lock bool //是否锁定

	player *Player //座位上的玩家
}

func (seat *Seat) SetPlayer(player *Player) {
	if seat.lock {
		exception.GameException(protocol.Code_invalidSeat)
	}
	player.RemoveRole(constant.RoleViewer)
	player.AddRole(constant.RolePlayer)
	player.Seat = seat.id
	seat.player = player

}

func (seat *Seat) RemovePlayer() *Player {
	player := seat.player
	if player != nil {
		player.Seat = 0
	}
	seat.player = nil
	return player
}

func (seat *Seat) UpdateLock(lock bool) {
	if lock && seat.player != nil {
		exception.GameException(protocol.Code_playerAlreadySeat)
	}
	seat.lock = lock
}

func (seat *Seat) GetPlayer() *Player {
	return seat.player
}

func (seat *Seat) CheckPlayer(playerID int64) bool {
	if seat.player == nil {
		return false
	}
	return seat.player.GetId() == playerID
}

//是否空闲
func (seat *Seat) IsFree() bool {
	return !seat.lock && seat.player == nil
}

func (seat *Seat) BuildProtocol() *protocol.Seat {
	 result := &protocol.Seat{Id:seat.id, Lock:seat.lock}
	 if seat.player != nil {
	 	result.Player = seat.player.Player
	 }
	 return result
}

