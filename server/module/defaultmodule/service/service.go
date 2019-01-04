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
	"github.com/KylinHe/aliensboot-core/chanrpc"
	"github.com/KylinHe/aliensboot-core/cluster/center"
	"github.com/KylinHe/aliensboot-core/cluster/center/service"
	"github.com/KylinHe/aliensboot-core/exception"
	"github.com/KylinHe/aliensboot-core/protocol/base"
	"github.com/KylinHe/aliensboot-server/module/defaultmodule/conf"
	"github.com/KylinHe/aliensboot-server/protocol"
	"github.com/gogo/protobuf/proto"
)

var instance service.IService = nil

func Init(chanRpc *chanrpc.Server) {
	instance = center.PublicService(conf.Config.Service, service.NewRpcHandler(chanRpc, handle))
}

func Close() {
	center.ReleaseService(instance)
}

func handle(request *base.Any) *base.Any {
	requestProxy := &protocol.Request{}
	responseProxy := &protocol.Response{}
	response := &base.Any{}
	defer func() {
		//处理消息异常
		if err := recover(); err != nil {
			switch err.(type) {
			case protocol.Code:
				responseProxy.Code = err.(protocol.Code)
				break
			default:
				exception.PrintStackDetail(err)
				responseProxy.Code = protocol.Code_ServerException
			}
		}
		data, _ := proto.Marshal(responseProxy)
		responseProxy.Session = requestProxy.GetSession()
		response.Value = data
	}()
	error := proto.Unmarshal(request.Value, requestProxy)
	if error != nil {
		exception.GameException(protocol.Code_InvalidRequest)
	}
	handleRequest(request.GetAuthId(), request.GetGateId(), requestProxy, responseProxy)
	return response
}

func handleRequest(authID int64, gateID string, request *protocol.Request, response *protocol.Response) {

	//TODO Gen handler

	response.Code = protocol.Code_InvalidRequest
}
