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
func handleUserLogin(request *protocol.C2S_UserLogin, response *protocol.S2C_UserLogin) int64 {
	username := request.GetUsername()
	password := request.GetPassword()

	userCache := cache.GetUser(username)
	if userCache == nil {
		passwordHash := PasswordHash(username, password)

		userCache = cache.NewUser(username, passwordHash, "", "", "", "")
		//result.Result = passport.LoginLoginRet_invalidUser
		//return
	}

	passwordHash := PasswordHash(username, password)
	//密码不对
	if passwordHash != userCache.Password {
		response.Result = protocol.LoginResult_invalidPwd
		return 0
	}

	//更新ip
	//qdoc := bson.M{"_id": userCache.ID}
	//udoc := bson.M{"$set": bson.M{"ip": getNetworkAddress(network)}}
	//db.DatabaseHandler.Update(userCache.Name(), qdoc, udoc)
	response.Uid = userCache.GetId()
	token := NewToken()
	cache.PassportCache.SetUserToken(userCache.GetId(), token)
	response.Token = token
	response.Result = protocol.LoginResult_loginSuccess
	return response.GetUid()
}
