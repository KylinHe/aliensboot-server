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
	"github.com/KylinHe/aliensboot-core/cluster/center/service"
	"github.com/KylinHe/aliensboot-server/constant"
	"github.com/KylinHe/aliensboot-server/dispatch/lpc"
	"github.com/KylinHe/aliensboot-server/module/passport/core"
	"github.com/KylinHe/aliensboot-server/protocol"
	"gopkg.in/mgo.v2/bson"
	"time"
)

//
func handleUserLogin(ctx *service.Context, request *protocol.UserLogin, response *protocol.UserLoginRet) {
	username := request.GetUsername()
	password := request.GetPassword()

	user := core.GetUserByUsername(username)
	if user == nil {
		user = core.NewUser(username, password, "", "", "", "")
	}
	ip := ctx.GetHeaderStr(constant.HeaderIP)
	response.Result = core.ValidateState(user, ip)
	if response.Result != protocol.LoginResult_loginSuccess {
		return
	}
	response.Uid = user.GetId()
	//response.Token = NewToken()
	//cache.PassportCache.SetUserToken(uid, response.Token)
	response.ServerTime = time.Now().UnixNano() / (1000 * 1000)
	ctx.Auth(user.GetId())
	lpc.LogHandler.AddLog(&protocol.LogLogin{
		Id:     bson.NewObjectId().Hex(),
		RoleId: user.GetId(),
		Ip:     ip,
		Time:   time.Now().Unix(),
	})
}
