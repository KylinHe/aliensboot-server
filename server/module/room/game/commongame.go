/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/11/13
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package game

import (
	"github.com/KylinHe/aliensboot-server/protocol"
)

type CommonGameFactory struct {
}

func (this *CommonGameFactory) Match(appID string) bool {
	return appID == "0"
}

func (this *CommonGameFactory) NewGame(handler Handler) Game {
	return &CommonGame{Handler: handler}
}

type CommonGame struct {
	Handler
	starting bool
}

func (game *CommonGame) IsStart() bool {
	return game.starting
}

//开始游戏
func (game *CommonGame) Start() {
	game.starting = true
	//通知所有玩家游戏开始
	push := &protocol.Response{Room: &protocol.Response_GameStartRet{GameStartRet: &protocol.GameStartRet{}}}
	game.BroadcastOtherPlayer(-1, push)
}

//结束游戏
func (game *CommonGame) Stop() {
	game.starting = false
	push := &protocol.Response{Room: &protocol.Response_GameResetRet{GameResetRet: &protocol.GameResetRet{}}}
	game.BroadcastOtherPlayer(-1, push)
}

//接收玩家数据，同步给其他玩家
func (game *CommonGame) AcceptPlayerData(playerID int64, data string) {
	//游戏结束不处理游戏内的数据转发
	//if !game.IsStart() {
	//	return
	//}
	push := &protocol.Response{Room: &protocol.Response_GameDataRet{
		GameDataRet: &protocol.GameDataRet{
			Data: data,
		},
	},
	}
	game.BroadcastOtherPlayer(playerID, push)
}

func (game *CommonGame) AcceptPlayerMessage(playerID int64, request interface{}, response interface{}) {

}
