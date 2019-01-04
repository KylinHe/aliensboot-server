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

import "github.com/gogo/protobuf/proto"

type Game interface {
	Start()                                                                        //启动游戏
	IsStart() bool                                                                 //是否启动游戏
	Stop()                                                                         //结束游戏
	AcceptPlayerData(playerID int64, data string)                                  //接收玩家数据
	AcceptPlayerMessage(playerID int64, request interface{}, response interface{}) //接收玩家发送的消息
}

type Factory interface {
	Match(appID string) bool
	NewGame(handler Handler) Game
}

type Handler interface {
	BroadcastOtherPlayer(playerID int64, message proto.Message) //广播其他玩家
}
