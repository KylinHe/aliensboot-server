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
	"github.com/KylinHe/aliensboot-server/module/game/core"
	"github.com/KylinHe/aliensboot-server/protocol"
)

//
func handleChangeNickname(authID int64, gateID string, request *protocol.ChangeNickname, response *protocol.ChangeNicknameRet) {
	userSession := core.UserManager.GetUser(authID)
	userSession.ChangeNickname(request.GetNewName())
	response.Result = true
}
