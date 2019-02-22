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
	"encoding/json"
	"github.com/KylinHe/aliensboot-core/exception"
	"github.com/KylinHe/aliensboot-core/log"
	"github.com/KylinHe/aliensboot-server/protocol"
)

type BigoGameFactory struct {
}

func (this *BigoGameFactory) NewGame(handler Handler) Game {
	return &BigoGame{CommonGame: &CommonGame{Handler: handler}, data: make(map[string]interface{})}
}

type BigoGame struct {
	*CommonGame

	data map[string]interface{} //游戏数据

	dataStr string //转换成字符串的游戏数据

	ts int64
}

/**
 * 向直播间输出游戏状态数据(V2, 增量或全量)  主播/嘉宾
 * @param stateData          游戏状态数据
 * @param type          数据类型 0 - 增量数据 1 - 全量数据
 * @param ts            时间戳 单位毫秒
 * @param forceUpdate   强制所有人更新这次全量数据
 */

func (game *BigoGame) AcceptPlayerMessage(playerID int64, request interface{}, response interface{}) {
	log.Debugf("accept %v - %v - %v", playerID, request, response)
	switch request.(type) {
	case *protocol.UpdateBigoData:
		game.handleUpdateBigoData(playerID, request.(*protocol.UpdateBigoData))
	case *protocol.GetBigoData:
		game.handleGetBigoData(playerID, request.(*protocol.GetBigoData), response.(*protocol.GetBigoDataRet))
	default: //类型为其他类型时执行
		//
	}

}

func (game *BigoGame) handleGetBigoData(playerID int64, request *protocol.GetBigoData, response *protocol.GetBigoDataRet) {
	response.Type = request.Type
	response.Data = game.dataStr
	response.Ts = game.ts

	log.Debugf("get bigo data %v", game.dataStr)
}

func (game *BigoGame) handleUpdateBigoData(playerID int64, data *protocol.UpdateBigoData) {
	game.updateData(data.GetData(), data.GetType(), data.GetTs())

	//是否通知其他玩家
	if data.GetForceUpdate() {
		push := &protocol.Response{Room: &protocol.Response_UpdateBigoDataRet{
			UpdateBigoDataRet: &protocol.UpdateBigoDataRet{
				Type: data.GetType(),
				Ts:   data.GetTs(),
				Data: game.dataStr,
			},
		},
		}
		game.BroadcastOtherPlayer(playerID, push)
	}
}

func (game *BigoGame) updateData(data string, updateType int32, ts int64) {
	newData := make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &newData)
	if err != nil {
		exception.GameException(protocol.Code_gameInvalidMsgFormat)
	}

	if updateType == 0 {
		game.appendData(newData)
	} else if updateType == 1 {
		game.data = newData
	}
	result, err := json.Marshal(game.data)
	if err != nil {
		log.Debugf("invalid bigo data : %v", err)
	}
	game.dataStr = string(result)
	game.ts = ts

	log.Debugf("update bigo data %v - %v", game.ts, game.dataStr)
}

func (game *BigoGame) appendData(newData map[string]interface{}) {
	for key, value := range newData {
		game.data[key] = value
	}
}
