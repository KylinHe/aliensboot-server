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
	"github.com/KylinHe/aliensboot-server/dispatch"
	"github.com/KylinHe/aliensboot-server/protocol"
	"github.com/KylinHe/aliensboot-core/common/util"
	"github.com/KylinHe/aliensboot-core/exception"
	"github.com/KylinHe/aliensboot-core/log"
)

type rpcHandler struct {
	name string
}

func (this *rpcHandler) Request(node string, request *protocol.Request) *protocol.Response {
	var rpcRet *protocol.Response = nil
	var err error = nil
	if node != "" {
		rpcRet, err = dispatch.RequestNodeMessage(this.name, node, request)
	} else {
		rpcRet, err = dispatch.RequestMessage(this.name, request, "")
	}
	if err != nil {
		log.Error(err)
		exception.GameException(protocol.Code_InvalidService)
	}
	return rpcRet
}

func (this *rpcHandler) TestRequest(authId int64, node string, request *protocol.Request) (*protocol.Response, error) {
	if node != "" {
		return dispatch.TestRequestNodeMessage(authId, this.name, node, request)
	} else {
		return dispatch.TestRequestMessage(authId, this.name, request, util.Int64ToString(authId))
	}
}

func (this *rpcHandler) Send(node string, request *protocol.Request) error {
	if node != "" {
		return dispatch.SendNodeMessage(this.name, node, request)
	} else {
		return dispatch.SendMessage(this.name, request, "")
	}
}
