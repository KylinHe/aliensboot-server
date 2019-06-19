/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/8/25
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package rpc

import (
	"github.com/KylinHe/aliensboot-core/cluster/center"
	"github.com/KylinHe/aliensboot-server/dispatch"
	"github.com/KylinHe/aliensboot-server/protocol"
	"github.com/gogo/protobuf/proto"
)

func (this *gateRPCHandler) BindService1(node string, authID int64, service string) error {
	binds := make(map[string]string)
	binds[service] = center.ClusterCenter.GetNodeID()
	request := &protocol.BindService{
		AuthID: authID,
		Binds:  binds,
	}
	return this.BindService(node, request)
}

//推送玩家消息
func (this *gateRPCHandler) Push(fromService string, authID int64, node string, response *protocol.Response) error {
	data, err := proto.Marshal(response)
	if err != nil {
		return err
	}
	pushMessage := &protocol.PushMessage{
		AuthID:  authID,
		Data:    data,
		Service: fromService,
	}
	return this.PushMessage(node, pushMessage)
}

func (this *gateRPCHandler) BroadcastAll(node string, response *protocol.Response) {
	data, _ := proto.Marshal(response)
	message := &protocol.Request{
		Gate: &protocol.Request_PushMessage{
			PushMessage: &protocol.PushMessage{
				AuthID: -1,
				Data:   data,
			},
		},
	}
	dispatch.Broadcast(this.name, message)
}
