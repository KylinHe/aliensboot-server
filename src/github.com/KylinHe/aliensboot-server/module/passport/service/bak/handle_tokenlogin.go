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
func handleTokenLogin(request *protocol.C2S_TokenLogin, response *protocol.S2C_TokenLogin) int64 {
	token, _ := cache.PassportCache.GetUserToken(request.GetUid())
	if token != request.GetToken() {
		response.Result = protocol.LoginResult_tokenExpire
		return 0
	}
	response.Result = protocol.LoginResult_loginSuccess
	return request.GetUid()
}
