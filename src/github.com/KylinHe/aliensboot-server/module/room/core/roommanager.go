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
	"github.com/KylinHe/aliensboot-server/module/room/game/bigo"
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
	RoomManager.RegisterGameFactory("1", &bigo.GameFactory{})
	RoomManager.RegisterGameFactory("2", &agar.GameFactory{})
}

type roomManager struct {
	gameFactories map[string]game.Factory //游戏工厂类

	rooms map[string]*Room //运行的游戏  游戏id - 房间对象

	players map[int64]string //所有玩家的对应信息 玩家id - 房间id
}

func (this *roomManager) RegisterGameFactory(appId string, factory game.Factory) {
	this.gameFactories[appId] = factory
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

//房主创建新房间 玩家需要主动加入房间
func (this *roomManager) CreateRoomByPlayer(appID string, playerID int64, roomID string, force bool) *Room {
	if roomID != "" && force {
		this.CloseRoom(playerID, roomID)
	}
	roomConfig := conf.GetRoomConfig(appID)
	if roomConfig == nil {
		exception.GameException(protocol.Code_appIDNotFound)
	}
	//不能房主管理
	if !roomConfig.Anchor {
		exception.GameException(protocol.Code_invalidAuth)
	}
	room := this.newRoom(roomConfig, roomID)
	room.AddPlayer(playerID, constant.RoleAnchor | constant.RolePlayer)
	//room.InitPlayers(players)
	this.rooms[room.GetID()] = room
	this.players[playerID] = room.GetID()
	return room
}

//系统创建房间 玩家自动加入房间
func (this *roomManager) CreateRoom(appID string, roomID string, force bool, playerIDs []int64) *Room {
	if roomID != "" && force {
		this.CloseRoom(-1, roomID)
	}
	roomConfig := conf.GetRoomConfig(appID)
	room := this.newRoom(roomConfig, roomID)

	for _, playerID := range playerIDs {
		room.AddPlayer(playerID, constant.RolePlayer)
		this.players[playerID] = room.GetID()
	}

	this.rooms[room.GetID()] = room


	return room
}

//关闭房间
func (this *roomManager) CloseRoom(authID int64, roomID string) bool {
	room := this.rooms[roomID]
	if room == nil {
		return false
	}
	//玩家关闭的房间需要验证权限、只有房主能关闭房间
	if authID > 0 {
		room.EnsureAnchor(authID)
	}

	room.Close(func(playerID int64) {
		delete(this.players, playerID)
	})
	delete(this.rooms, roomID)
	return true
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
		viewers: make(Viewers),
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
