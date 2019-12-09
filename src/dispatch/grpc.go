/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/8/22
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package dispatch

import (
	"github.com/KylinHe/aliensboot-core/cluster/message"
	"github.com/KylinHe/aliensboot-core/protocol/base"
	"github.com/KylinHe/aliensboot-server/constant"
	"github.com/KylinHe/aliensboot-server/protocol"
	"github.com/gogo/protobuf/proto"
)

var serviceMapping = make(map[string]*message.RemoteService)

// 初始化远程服务
func Init(service string) {
	allocService(service)
}

////阻塞请求消息 - 根据负载均衡动态分配一个节点处理
//func RequestMessage(serviceName string, message *protocol.Request, hashKey string) (*protocol.Response, error) {
//	data, err := proto.Marshal(message)
//	if err != nil {
//		return nil, err
//	}
//	request := &base.Any{AuthId: constant.SystemAuthId, Value: data}
//	response, err := Request(serviceName, request, hashKey)
//	if err != nil {
//		return nil, err
//	}
//	messageRet := &protocol.Response{}
//	err = messageRet.Unmarshal(response.GetValue())
//	if err != nil {
//		return nil, err
//	}
//	return messageRet, nil
//}

func RequestMessage(authId int64, serviceName string, hashKey string, message *protocol.Request, header map[string][]byte) (*protocol.Response, error) {
	data, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}
	request := &base.Any{AuthId: authId, Value: data, Header: header}
	response, err := Request(serviceName, hashKey, request)
	if err != nil {
		return nil, err
	}
	messageRet := &protocol.Response{}
	err = messageRet.Unmarshal(response.GetValue())
	if err != nil {
		return nil, err
	}
	return messageRet, nil
}

//同步阻塞请求
func RequestNodeMessage(authId int64, serviceName string, serviceID string, message *protocol.Request, header map[string][]byte) (*protocol.Response, error) {
	data, _ := message.Marshal()
	request := &base.Any{AuthId: authId, Value: data, Header: header}
	response, err := RequestNode(serviceName, serviceID, request)
	if err != nil {
		return nil, err
	}
	messageRet := &protocol.Response{}
	_ = messageRet.Unmarshal(response.GetValue())
	return messageRet, nil
}

//异步发送信息
func SendMessage(serviceName string, message *protocol.Request, hashKey string, header map[string][]byte) error {
	data, _ := message.Marshal()
	request := &base.Any{AuthId: constant.SystemAuthId, Value: data, Header: header}
	return Send(serviceName, request, hashKey)
}

//异步发送信息
func SendNodeMessage(serviceName string, serviceID string, message *protocol.Request, header map[string][]byte) error {
	data, _ := message.Marshal()
	request := &base.Any{AuthId: constant.SystemAuthId, Value: data, Header: header}
	return SendNode(serviceName, serviceID, request)
}

func Broadcast(serviceName string, message *protocol.Request, header map[string][]byte) {
	service := allocService(serviceName)
	data, _ := message.Marshal()
	request := &base.Any{AuthId: constant.SystemAuthId, Value: data, Header: header}
	service.BroadcastAll(request)
}

func Request(serviceName string, hashKey string, message *base.Any) (*base.Any, error) {
	service := allocService(serviceName)
	return service.Request(message, hashKey)
}

func RequestNode(serviceName string, serviceID string, message *base.Any) (*base.Any, error) {
	service := allocService(serviceName)
	return service.RequestNode(serviceID, message)
}

//func AsyncRequest(serviceName string, hashKey string, asyncCall *service.AsyncCall) error {
//	service := allocService(serviceName)
//	return service.AsyncRequest(hashKey, asyncCall)
//}
//
//func AsyncRequestNode(serviceName string, serviceID string, asyncCall *service.AsyncCall) error {
//	service := allocService(serviceName)
//	return service.AsyncRequestNode(serviceID, asyncCall)
//}

func Send(serviceName string, message *base.Any, hashKey string) error {
	service := allocService(serviceName)
	return service.Send(message, hashKey)
}

func SendNode(serviceName string, serviceID string, message *base.Any) error {
	service := allocService(serviceName)
	return service.SendNode(serviceID, message)
}

func allocService(serviceName string) *message.RemoteService {
	service := serviceMapping[serviceName]
	if service == nil {
		service = message.NewRemoteService(serviceName)
		serviceMapping[serviceName] = service
	}
	return service
}
