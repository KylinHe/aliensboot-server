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
	"github.com/KylinHe/aliensboot-core/protocol/base"
	"github.com/KylinHe/aliensboot-server/module/gate/network"
	"github.com/KylinHe/aliensboot-server/module/gate/route"
	"github.com/KylinHe/aliensboot-server/protocol"
)

//
func handlePushMessage(request *protocol.PushMessage) {
	msgID := route.GetServiceSeq(request.GetService())
	msg := &base.Any{Id: msgID, Value: request.GetData()}
	authID := request.GetAuthID()
	if authID == -1 {
		network.Manager.Broadcast(msg)
	} else if authID > 0 {
		network.Manager.Push(authID, msg)
	}
}
