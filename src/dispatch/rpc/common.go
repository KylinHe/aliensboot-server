/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/11/14
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package rpc

import (
	"github.com/KylinHe/aliensboot-core/exception"
	"github.com/KylinHe/aliensboot-core/log"
	"github.com/KylinHe/aliensboot-server/constant"
	"github.com/KylinHe/aliensboot-server/dispatch"
	"github.com/KylinHe/aliensboot-server/protocol"
)

type rpcHandler struct {
	name string
}

func EnsureCode(response *protocol.Response, err error) (*protocol.Response, error) {
	code := response.Code
	if err != nil {
		log.Error(err)
		code = protocol.Code_InvalidService
	}
	if code != protocol.Code_Success {
		exception.GameException(code)
	}
	return response, err
}

// 根据负载均衡策略发送请求
func (this *rpcHandler) AuthRequest(authId int64, hashKey string, request *protocol.Request) (*protocol.Response, error) {
	return dispatch.RequestMessage(authId, this.name, hashKey, request, nil)
}

func (this *rpcHandler) AuthNodeRequest(authId int64, node string, request *protocol.Request) (*protocol.Response, error) {
	return dispatch.RequestNodeMessage(authId, this.name, node, request, nil)
}

// 向指定节点发送请求
func (this *rpcHandler) Request(node string, request *protocol.Request) (*protocol.Response, error) {
	if node != "" {
		return EnsureCode(dispatch.RequestNodeMessage(constant.SystemAuthId, this.name, node, request, nil))
	} else {
		return EnsureCode(dispatch.RequestMessage(constant.SystemAuthId, this.name, "", request, nil))
	}
}

func (this *rpcHandler) Send(node string, request *protocol.Request) error {
	if node != "" {
		return dispatch.SendNodeMessage(this.name, node, request, nil)
	} else {
		return dispatch.SendMessage(this.name, request, "", nil)
	}
}

// 广播请求
func (this *rpcHandler) Broadcast(request *protocol.Request) {
	dispatch.Broadcast(this.name, request, nil)
}
