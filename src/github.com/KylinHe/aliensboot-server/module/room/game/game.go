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
	AcceptPlayerData(playerID int64, data string, roles int32)                                  //接收玩家数据
}

type Factory interface {
	NewGame(handler Handler) Game
}

type Handler interface {
	BroadcastOtherPlayer(playerID int64, roles int32, message proto.Message) //广播其他玩家
}
