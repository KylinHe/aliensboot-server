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
	"github.com/KylinHe/aliensboot-core/common/util"
	"github.com/KylinHe/aliensboot-core/exception"
	"github.com/KylinHe/aliensboot-server/constant"
	"github.com/KylinHe/aliensboot-server/module/room/config"
	"github.com/KylinHe/aliensboot-server/module/room/game"
	"github.com/KylinHe/aliensboot-server/protocol"
	"github.com/gogo/protobuf/proto"
)

type Room struct {
	id string //房间id

	Seats //桌子,数组下标为座位编号

	config *config.RoomConfig //房间配置

	game game.Game //房间内进行的游戏对象

	state int32 //0-准备中，1-游戏中，2-游戏房间已结束，3-回合结束

	viewers Viewers

	timerMgr *util.TimerManager
}

func (room *Room) GetID() string {
	return room.id
}

func (room *Room) UpdateSeat(authID int64, seatID int32, opt int32) {
	player := room.Get(authID)
	if player == nil {
		exception.GameException(protocol.Code_invalidAuth)
	}

	seat := room.EnsureSeat(seatID)
	var oldSeat *Seat = nil

	switch opt {
	case constant.OptLockSeat:
		if !player.IsAnchor() {
			exception.GameException(protocol.Code_invalidAuth)
		}
		seat.UpdateLock(true)
		break
	case constant.OptUnlockSeat:
		if !player.IsAnchor() {
			exception.GameException(protocol.Code_invalidAuth)
		}
		seat.UpdateLock(false)
		break
	case constant.OptLeaveSeat:
		if !player.IsAnchor() && !seat.CheckPlayer(authID) {
			exception.GameException(protocol.Code_invalidAuth)
		}
		player := seat.RemovePlayer()
		room.AddViewer(player)
		break
	case constant.OptChangeSeat:
		//位置没有变化
		if player.GetSeat() == seatID {
			return
		}
		if !seat.IsFree() {
			exception.GameException(protocol.Code_invalidSeat)
		}

		oldSeat = room.EnsureSeat(player.GetSeat())
		player := oldSeat.RemovePlayer()
		seat.SetPlayer(player)
		break
	}

	seatData := []*protocol.Seat{seat.BuildProtocol()}
	if oldSeat != nil {
		seatData = append(seatData, oldSeat.BuildProtocol())
	}
	push := &protocol.Response{Room: &protocol.Response_UpdateSeatPush{UpdateSeatPush: &protocol.UpdateSeatPush{
		RoomID:room.GetID(),
		Seats:seatData,
	}}}

	room.BroadcastOtherPlayer(-1, constant.RoleAll, push)
}


func (room *Room) AddViewer(player *Player) {
	if !room.config.Viewer {
		exception.GameException(protocol.Code_invalidAuth)
	}
	room.viewers.AddViewer(player)
}


func (room *Room) BuildProtocol() *protocol.Room {
	return &protocol.Room{
		RoomID:room.GetID(),
		Seats:room.GetAllSeatData(),
	}
}


func (room *Room) JoinSeat(authorId int64, acceptID int64, seatID int32) {
	room.EnsureAnchor(authorId)

	player := room.viewers.GetViewer(acceptID)
	if player == nil {
		exception.GameException(protocol.Code_playerNotFound)
	}
	seat := room.Add(player, seatID)
	room.viewers.RemoveViewer(acceptID)

	push := &protocol.Response{Room: &protocol.Response_UpdateSeatPush{UpdateSeatPush: &protocol.UpdateSeatPush{
		RoomID:room.GetID(),
		Seats: []*protocol.Seat{seat.BuildProtocol()},
	}}}

	//通知所有玩家
	room.BroadcastOtherPlayer(-1, constant.RoleAll, push)
}


func (room *Room) RequestJoinGame(request *protocol.JoinRequest) {
	player := room.viewers.GetViewer(request.GetPlayerID())
	if player == nil {
		exception.GameException(protocol.Code_playerNotFound)
	}
	//转发给房主
	push := &protocol.Response{Room: &protocol.Response_RequestJoinSeatPush{RequestJoinSeatPush: &protocol.RequestJoinSeatPush{
		Request:request,
	}}}
	room.SendToAnchor(push)
}

//新增玩家
func (room *Room) AddPlayer(playerID int64, roles int32) *Player {
	player := &protocol.Player{
		Id: playerID,
		Nickname: "玩家" + util.Int64ToString(playerID),
		Role: roles,
	}
	result := &Player{Player: player}
	if result.HaveRole(constant.RoleViewer) {
		room.AddViewer(result)
	} else {
		room.Add(result, constant.AnySeat)
	}
	return result
}


//关闭房间
func (room *Room) Close(callback func(playerID int64)) {
	//TODO 通知T人
	room.GameOver()

	push := &protocol.Response{Room: &protocol.Response_RoomClosePush{RoomClosePush: &protocol.RoomClosePush{RoomID:room.GetID()}}}
	room.BroadcastOtherPlayer(-1, constant.RoleAll, push)
	room.Foreach(func(player *Player) {
		callback(player.GetId())
	})
	room.viewers.ForeachViewer(func(viewer *Player) {
		callback(viewer.GetId())
	})
	room.Seats.Clean()
}

func (room *Room) GameOver() {
	room.game.Stop()
}

//获取所有玩家数据
func (room *Room) GetAllSeatData() []*protocol.Seat {
	results := []*protocol.Seat{}
	room.ForeachSeat(func(seat *Seat) {
		results = append(results, seat.BuildProtocol())
	})
	return results
}

func (room *Room) GetPlayerData(playerID int64) *protocol.Player {
	player := room.EnsurePlayer(playerID)
	return player.Player
}

func (room *Room) EnsurePlayer(playerID int64) *Player {
	player := room.Get(playerID)
	if player == nil {
		exception.GameException(protocol.Code_playerNotFound)
	}
	return player
}

func (room *Room) EnsureAnchor(playerID int64) *Player {
	player := room.EnsurePlayer(playerID)
	if !player.HaveRole(constant.RoleAnchor) {
		exception.GameException(protocol.Code_invalidAuth)
	}
	return player
}

//玩家准备
func (room *Room) PlayerReady(playerID int64) {
	if room.IsGameStart() {
		exception.GameException(protocol.Code_gameAlreadyStart)
	}
	player := room.EnsurePlayer(playerID)
	player.Ready()

	//所有玩家准备完毕、即可开始游戏
	if room.IsAllReady() {
		//启动新游戏
		room.game.Start()
	}
}

func (room *Room) GameStart(authID int64) {
	room.EnsureAnchor(authID)
	if room.IsGameStart() {
		exception.GameException(protocol.Code_gameAlreadyStart)
	}
	//启动新游戏
	room.game.Start()
}

//游戏是否开始
func (room *Room) IsGameStart() bool {
	return room.game.IsStart()
}

//广播其他玩家
func (room *Room) BroadcastOtherPlayer(playerID int64, roles int32, message proto.Message) {
	sendData, _ := proto.Marshal(message)
	room.Foreach(func(player *Player) {
		if player.GetId() != playerID && player.HaveRole(roles) {
			player.SendMsg(sendData)
		}
	})

	if roles & constant.RoleViewer != 0 {
		room.viewers.ForeachViewer(func(viewer *Player) {
			viewer.SendMsg(sendData)
		})
	}
}

func (room *Room) SendToPlayer(playerID int64, roles int32, message proto.Message) {
	sendData, _ := proto.Marshal(message)
	room.Foreach(func(player *Player) {
		if player.GetId() == playerID && player.HaveRole(roles) {
			player.SendMsg(sendData)
		}
	})
}

func (room *Room) GetTimerMgr() *util.TimerManager {
	return room.timerMgr
}

func (room *Room) SendToAnchor(message proto.Message) {
	sendData, _ := proto.Marshal(message)

	room.Foreach(func(player *Player) {
		if player.HaveRole(constant.RoleAnchor) {
			player.SendMsg(sendData)
		}
	})
}

//接收玩家数据，同步给其他玩家
func (room *Room) AcceptPlayerData(playerID int64, data string, roles int32) {
	room.game.AcceptPlayerData(playerID, data, roles)
}
