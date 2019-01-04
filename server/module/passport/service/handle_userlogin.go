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
	"github.com/KylinHe/aliensboot-server/module/passport/cache"
	"github.com/KylinHe/aliensboot-server/protocol"
)

//
func handleUserLogin(request *protocol.UserLogin, response *protocol.UserLoginRet) int64 {
	username := request.GetUsername()
	password := request.GetPassword()
	passwordHash := PasswordHash(username, password)

	uid, _ := cache.PassportCache.GetUidByUsername(username)

	if uid <= 0 {
		uid = cache.NewUser(username, passwordHash, "", "", "", "").GetId()
	}

	response.Uid = uid
	//response.Token = NewToken()
	//cache.PassportCache.SetUserToken(uid, response.Token)
	response.Result = protocol.LoginResult_loginSuccess
	return response.Uid
}
