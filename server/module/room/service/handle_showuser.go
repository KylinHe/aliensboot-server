/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/3/30
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package service

import (
	"github.com/KylinHe/aliensboot-server/module/room/core"
	"github.com/KylinHe/aliensboot-server/protocol"
)

//
func handleShowUser(authID int64, gateID string, request *protocol.ShowUser, response *protocol.ShowUserRet) {
	game := core.RoomManager.GetRoomByPlayerID(authID)
	response.Player = game.GetPlayerData(request.GetPlayerid())
}
