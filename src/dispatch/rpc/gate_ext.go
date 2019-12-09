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
	"github.com/KylinHe/aliensboot-server/protocol"
	"github.com/gogo/protobuf/proto"
)

func (this *gateRPCHandler) BindService1(node string, authId int64, service string) bool {
	binds := make(map[string]string)
	binds[service] = center.ClusterCenter.GetNodeID()
	request := &protocol.BindService{
		AuthId: authId,
		Binds:  binds,
	}
	return this.BindService(node, request).GetResult()
}

// 推送玩家相关信息
func (this *gateRPCHandler) PushRole(fromService string, authId int64, response *protocol.Response) error {
	return this.PushMultiRole(fromService, []int64{authId}, response)
}

func (this *gateRPCHandler) PushMultiRole(fromService string, authIds []int64, response *protocol.Response) error {
	// 玩家在线需要推送
	data, err := proto.Marshal(response)
	if err != nil {
		return err
	}
	pushMessage := &protocol.PushMessage{
		AuthId:  authIds,
		Data:    data,
		Service: fromService,
		Relay:   true, //允许转发
	}
	return this.PushMessage("", pushMessage)
}

//推送玩家消息到指定节点
func (this *gateRPCHandler) PushNode(fromService string, authId int64, node string, response *protocol.Response) error {
	return this.PushNodeMulti(fromService, []int64{authId}, node, response)
}

//推送多个玩家消息到指定节点
func (this *gateRPCHandler) PushNodeMulti(fromService string, authIds []int64, node string, response *protocol.Response) error {
	data, err := proto.Marshal(response)
	if err != nil {
		return err
	}
	pushMessage := &protocol.PushMessage{
		AuthId:  authIds,
		Data:    data,
		Service: fromService,
	}
	return this.PushMessage(node, pushMessage)
}

//广播所有用户
func (this *gateRPCHandler) BroadcastAll(fromService string, response *protocol.Response) {
	data, _ := proto.Marshal(response)
	message := &protocol.Request{
		Gate: &protocol.Request_PushMessage{
			PushMessage: &protocol.PushMessage{
				All:     true,
				Data:    data,
				Service: fromService,
			},
		},
	}
	this.Broadcast(message)
}

//广播多个玩家
func (this *gateRPCHandler) BroadcastMulti(fromService string, authId []int64, response *protocol.Response) {
	data, _ := proto.Marshal(response)
	message := &protocol.Request{
		Gate: &protocol.Request_PushMessage{
			PushMessage: &protocol.PushMessage{
				All:     false,
				AuthId:  authId,
				Data:    data,
				Service: fromService,
			},
		},
	}
	this.Broadcast(message)
}
