/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2019/02/15
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package agar

import (
	"github.com/KylinHe/aliensboot-core/log"
	"github.com/KylinHe/aliensboot-server/constant"
	"github.com/KylinHe/aliensboot-server/module/room/game"
	"github.com/KylinHe/aliensboot-server/module/room/game/agar/collision"
	"github.com/KylinHe/aliensboot-server/module/room/game/agar/util"
	"github.com/KylinHe/aliensboot-server/protocol"
)

type GameFactory struct {

}

func (this *GameFactory) NewGame(handler game.Handler) game.Game {
	return &Game{CommonGame: &game.CommonGame{Handler: handler}}
}

type Game struct {
	*game.CommonGame

	battleId 	string
	users	 	map[int32]*BattleUser
	tickCount 	int32
	gameOverTick int32
	lastSysTick  int32

	mapBorder    *util.MapBorder

	colMgr collision.ICollision
	ballIDCounter int32 //球的id计数器
}

//接收玩家数据，同步给其他玩家
func (game *Game) AcceptPlayerData(playerID int64, data string, roles int32)  {
	//游戏结束不处理游戏内的数据转发
	log.Debugf("accpet msg %v - %v", playerID, data)
	push := &protocol.Response{Room: &protocol.Response_GameDataPush {
			GameDataPush: &protocol.GameDataPush{
				Data: data,
			},
		},
	}
	game.BroadcastOtherPlayer(-1, constant.RoleAll, push)
}

/**
 * 向直播间输出游戏状态数据(V2, 增量或全量)  主播/嘉宾
 * @param stateData          游戏状态数据
 * @param type          数据类型 0 - 增量数据 1 - 全量数据
 * @param ts            时间戳 单位毫秒
 * @param forceUpdate   强制所有人更新这次全量数据
 */

//func (game *Game) AcceptPlayerMessage(playerID int64, request interface{}, response interface{}) {
//	log.Debugf("accept %v - %v - %v", playerID, request, response)
//
//
//}

