/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/6/8
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package core

import (
	"github.com/KylinHe/aliensboot-core/common/util"
	"github.com/KylinHe/aliensboot-core/exception"
	"github.com/KylinHe/aliensboot-server/constant"
	"github.com/KylinHe/aliensboot-server/module/room/conf"
	"github.com/KylinHe/aliensboot-server/module/room/config"
	"github.com/KylinHe/aliensboot-server/module/room/game"
	"github.com/KylinHe/aliensboot-server/module/room/game/agar"
	"github.com/KylinHe/aliensboot-server/protocol"
)

var RoomManager = &roomManager{
	gameFactories: make(map[string]game.Factory),
	rooms:         make(map[string]*Room),
	players:       make(map[int64]string),
}

func init() {
	//
	RoomManager.RegisterGameFactory("0", &game.CommonGameFactory{})
	RoomManager.RegisterGameFactory("1", &game.BigoGameFactory{})
	RoomManager.RegisterGameFactory("2", &agar.AgarGameFactory{})
}

type roomManager struct {
	gameFactories map[string]game.Factory //游戏工厂类

	rooms map[string]*Room //运行的游戏  游戏id - 房间对象

	players map[int64]string //所有玩家的对应信息 玩家id - 房间id

}

func (this *roomManager) RegisterGameFactory(appId string, factory game.Factory) {
	this.gameFactories[appId] = factory
}

func (this *roomManager) ChangePlayerState(authID int64, playerID int64, state int32) int32 {
	room := this.GetRoomByPlayerID(authID)
	auth := room.Get(authID)
	if !auth.IsAnchor() {
		return 1
	}
	if state == PlayerStatekick {
		room.kickPlayer(playerID)
		return 0
	}
	return 1
}

func (this *roomManager) ChangeGameState(authID int64, state int32) int32 {
	room := this.GetRoomByPlayerID(authID)
	auth := room.Get(authID)
	if !auth.IsAnchor() {
		return 1
	}

	if state == RoomStateOver {
		//通知所有玩家游戏结束
		room.game.Stop()
	}
	return 0
}

//获取玩家在哪个房间
func (this *roomManager) GetRoomByPlayerID(playerID int64) *Room {
	roomID := this.players[playerID]
	if roomID == "" {
		exception.GameException(protocol.Code_roomNotFound)
	}
	return this.EnsureRoom(roomID)
}

//获取房间
func (this *roomManager) EnsureRoom(roomID string) *Room {
	game := this.rooms[roomID]
	if game == nil {
		exception.GameException(protocol.Code_roomNotFound)
	}
	return game
}

//玩家加入房间
func (this *roomManager) JoinRoom(appID string, roomID string, playerID int64) *Room {
	room := this.EnsureRoom(roomID)
	room.AddPlayer(playerID, constant.RoleViewer)
	this.players[playerID] = room.GetID()
	return room
}

//玩家申请从观众成为主播
//func (this *roomManager) RequestJoinGame(playerID int64) *Room {
//	room := this.GetRoomByPlayerID(playerID)
//	room.RequestJoinGame(playerID)
//	return room
//}

//房主创建新房间
func (this *roomManager) CreateRoom(appID string, playerID int64, roomID string, force bool, maxSeat int32) *Room {
	if roomID != "" && force {
		this.RemoveRoom(roomID)
	}
	roomConfig := conf.GetRoomConfig(appID)

	if maxSeat > 0 {
		roomConfig = &config.RoomConfig{
			AppID:   roomConfig.AppID,
			MaxSeat: int(maxSeat),
		}
	}
	room := this.newRoom(roomConfig, roomID)
	//room.InitPlayers(players)
	this.rooms[room.GetID()] = room
	room.AddPlayer(playerID, constant.RoleAnchor)
	this.players[playerID] = room.GetID()
	//if players != nil {
	//	for _, player := range players {
	//		this.players[player.GetPlayerid()] = room.GetID()
	//	}
	//}
	return room
}

//关闭房间
func (this *roomManager) RemoveRoom(roomID string) {
	room := this.rooms[roomID]
	if room != nil {
		room.Close(func(playerID int64) {
			delete(this.players, playerID)
		})
		delete(this.rooms, roomID)
	}
}

//新建房间
func (this *roomManager) newRoom(config *config.RoomConfig, roomID string) *Room {
	if roomID != "" && this.rooms[roomID] != nil {
		exception.GameException(protocol.Code_roomAlreadyExist)
	}

	result := &Room{
		id:     roomID,
		config: config,
		Seats:  NewSeats(config.MaxSeat),
		viewer: make(map[int64]*Player),
	}

	if result.id == "" {
		result.id = util.GenUUID()
	}

	for appId, factory := range this.gameFactories {
		if appId == config.AppID {
			result.game = factory.NewGame(result)
			break
		}
	}

	if result.game == nil {
		exception.GameException(protocol.Code_appIDNotFound)
	}
	return result
}
