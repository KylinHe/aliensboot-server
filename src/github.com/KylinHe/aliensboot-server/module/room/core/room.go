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
	"github.com/KylinHe/aliensboot-core/log"
	"github.com/KylinHe/aliensboot-server/constant"
	"github.com/KylinHe/aliensboot-server/module/room/config"
	"github.com/KylinHe/aliensboot-server/module/room/game"
	"github.com/KylinHe/aliensboot-server/protocol"
	"github.com/gogo/protobuf/proto"
)

type Seat int32

const (
	RoomStateReady     int32 = 0
	RoomState          int32 = 1
	RoomStateOver      int32 = 2
	RoomStateRoundOver int32 = 3

	PlayerStateJoin int32 = 0
	PlayerStatekick int32 = 1
)

type Room struct {
	id string //房间id

	Seats //桌子,数组下标为座位编号

	config *config.RoomConfig //房间配置

	game game.Game //房间内进行的游戏对象

	state int32 //0-准备中，1-游戏中，2-游戏房间已结束，3-回合结束

	viewer map[int64]*Player //观众

}

func (room *Room) GetID() string {
	return room.id
}

func (room *Room) AcceptJoinGame(author int64, acceptID int64) {
	player := room.viewer[acceptID]
	if player == nil {
		exception.GameException(protocol.Code_playerNotFound)
	}

	//成为嘉宾
	player.GroupId = constant.RoleGuest

	//通知关注加入游戏
	push := &protocol.Response{Room: &protocol.Response_PlayerJoinRet{PlayerJoinRet: &protocol.PlayerJoinRet{
		RoomID: room.GetID(),
		Player: player.Player,
	}}}
	player.SendProtoMsg(push)

}

func (room *Room) JoinGame(playerID int64) {
	player := room.viewer[playerID]
	if player == nil || player.GroupId != constant.RoleGuest {
		exception.GameException(protocol.Code_playerNotFound)
	}

	room.AddPlayerToGame(player)
	delete(room.viewer, playerID)

	//通知主播加入游戏
	push := &protocol.Response{Room: &protocol.Response_PreJoinGameReq{PreJoinGameReq: &protocol.PreJoinGameReq{
		Player: player.Player,
	}}}
	room.SendToAnchor(push)
}

func (room *Room) GetViewer(playerID int64) *Player {
	return room.viewer[playerID]
}

func (room *Room) RequestJoinGame(playerID int64) {
	player := room.viewer[playerID]

	log.Debugf("request join game %v - %v : %v", playerID, player, room.GetAllPlayerData())
	if player == nil {
		exception.GameException(protocol.Code_playerNotFound)
	}
	push := &protocol.Response{Room: &protocol.Response_ContinueJoinGameReq{ContinueJoinGameReq: &protocol.ContinueJoinGameReq{
		PlayerID: playerID,
	}}}
	room.SendToAnchor(push)
}

//新增玩家
func (room *Room) AddPlayer(playerID int64, groupID int32) *Player {
	player := &protocol.Player{
		Playerid: playerID,
		Nickname: "蛇皮" + util.Int64ToString(playerID),
		GroupId:  groupID,
	}
	result := &Player{Player: player}
	if groupID == constant.RoleViewer {
		room.viewer[playerID] = result
	} else {
		room.AddPlayerToGame(result)
	}
	return result
}

func (room *Room) AddPlayerToGame(player *Player) {
	ok := room.Add(player)
	if !ok {
		exception.GameException(protocol.Code_roomMaxPlayer)
	}

	//room.BroadcastOtherPlayer(-1, push)
}

//一次性添加房间人员
//func (room *Room) InitPlayers(players []*protocol.Player) {
//	if players == nil {
//		return
//	}
//	//初始化玩家
//	for _, player := range players {
//		room.AddPlayer(player)
//	}
//}

//关闭房间
func (room *Room) Close(callback func(playerID int64)) {
	room.Foreach(func(player *Player) {
		player.kick(protocol.KickType_KickOut)
		callback(player.GetPlayerid())
	})
	room.Clean()
	if room.game != nil {
		room.game.Stop()
	}
}

//获取所有玩家数据
func (room *Room) GetAllPlayerData() []*protocol.Player {
	results := []*protocol.Player{}
	room.Foreach(func(player *Player) {
		results = append(results, player.Player)
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

//玩家上报游戏结果
//玩家上报结果
func (room *Room) UploadResult(playerID int64, reports []*protocol.PlayerResult) {
	//TODO 处理玩家的上报结果
	//game := room.EnsureGame()

}

//游戏是否开始
func (room *Room) IsGameStart() bool {
	return room.game != nil && room.game.IsStart()
}

func (room *Room) EnsureGame() game.Game {
	if room.game == nil {
		exception.GameException(protocol.Code_gameNotFound)
	}
	return room.game
}

//T人
func (room *Room) kickPlayer(playerID int64) *Player {
	player := room.Delete(playerID)
	if player != nil {
		player.kick(protocol.KickType_KickOut)
	}
	return player
}

//广播其他玩家
func (room *Room) BroadcastOtherPlayer(playerID int64, message proto.Message) {
	sendData, _ := proto.Marshal(message)
	room.Foreach(func(player *Player) {
		if player.GetPlayerid() != playerID {
			player.SendMsg(sendData)
		}
	})
}

//广播给所有观众
func (room *Room) BroadcastViewer(message proto.Message) {
	sendData, _ := proto.Marshal(message)
	for _, player := range room.viewer {
		player.SendMsg(sendData)
	}
}

func (room *Room) SendToAnchor(message proto.Message) {
	sendData, _ := proto.Marshal(message)
	room.Foreach(func(player *Player) {
		if player.GetGroupId() == constant.RoleAnchor {
			player.SendMsg(sendData)
		}
	})
}

//接收玩家数据，同步给其他玩家
func (room *Room) AcceptPlayerData(playerID int64, data string) {
	room.game.AcceptPlayerData(playerID, data)
}

//接收玩家消息
func (room *Room) AcceptPlayerMessage(playerID int64, request interface{}, response interface{}) {
	room.game.AcceptPlayerMessage(playerID, request, response)
}
