/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/3/30
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package bak

import (
	"github.com/KylinHe/aliensboot-server/module/passport/cache"
	"github.com/KylinHe/aliensboot-server/protocol"
)

//
func handleUserRegister(request *protocol.C2S_UserRegister, response *protocol.S2C_UserRegister) int64 {
	username := request.GetUsername()
	passwd := request.GetPassword()
	if cache.PassportCache.IsUsernameExist(username) {
		response.Msg = "user already exists"
		response.Result = protocol.RegisterResult_userExists
		return 0
	}

	passwd = PasswordHash(username, passwd)
	//TODO 有风险最好查询 数据库再加一层判断
	userCache := cache.NewUser(username, passwd, "", "", "", "")

	response.Result = protocol.RegisterResult_registerSuccess
	response.Uid = userCache.Id
	token := NewToken()
	cache.PassportCache.SetUserToken(userCache.Id, token)
	response.Token = token
	return response.GetUid()
}
