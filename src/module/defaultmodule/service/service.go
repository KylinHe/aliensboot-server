// Code generated by aliensboot. DO NOT EDIT.
// source: module defaultmodule
package service

import (
    "github.com/KylinHe/aliensboot-core/chanrpc"
    "github.com/KylinHe/aliensboot-core/cluster/center/service"
    "github.com/KylinHe/aliensboot-core/cluster/center"
    "github.com/KylinHe/aliensboot-core/exception"
    "github.com/KylinHe/aliensboot-core/protocol/base"
    "github.com/KylinHe/aliensboot-server/module/defaultmodule/conf"
    "github.com/KylinHe/aliensboot-server/protocol"
    "github.com/gogo/protobuf/proto"
)

var instance service.IService = nil

var handlers = make(map[uint16]func(request *base.Any)*base.Any)

func Init(chanRpc *chanrpc.Server) {
	instance = center.PublicService(conf.Config.Service, service.NewRpcHandler(chanRpc, handle))
}

func Close() {
	center.ReleaseService(instance)
}


//register self handler
func RegisterHandler(msgID uint16, handler func(request *base.Any)*base.Any) {
	handlers[msgID] = handler
}

func handleInternal(request *base.Any) (bool, *base.Any) {
	handler := handlers[request.Id]
	if handler == nil {
		return false, nil
	}
	response := handler(request)
	return true, response
}

func handle(ctx *service.Context) {
    ok, _ := handleInternal(ctx.Request)
    if ok {
    	return
    }
	requestProxy := &protocol.Request{}
	responseProxy := &protocol.Response{}
	isResponse := false
	defer func() {
		//处理消息异常
		if err := recover(); err != nil {
			switch err.(type) {
			case protocol.Code:
				responseProxy.Code = err.(protocol.Code)
				break
			case *protocol.CodeMessage:
			    responseProxy.CodeMessage = err.(*protocol.CodeMessage)
			    break
			default:
				exception.PrintStackDetail(err)
				responseProxy.Code = protocol.Code_ServerException
			}
			isResponse = true
		}
		if isResponse {
            _ = ctx.GOGOProto(responseProxy)
        }
	}()
	error := proto.Unmarshal(ctx.Request.Value, requestProxy)
	if error != nil {
		exception.GameException(protocol.Code_InvalidRequest)
	}
	responseProxy.Session = requestProxy.GetSession()
	isResponse = handleRequest(ctx.Request.GetAuthId(), ctx.Request.GetGateId(), requestProxy, responseProxy)
}

func handleRequest(authID int64, gateID string, request *protocol.Request, response *protocol.Response) bool {
	
	if request.GetBenchmark() != nil {
		messageRet := &protocol.BenchmarkRet{}
		handleBenchmark(authID, gateID, request.GetBenchmark(), messageRet)
		response.Defaultmodule = &protocol.Response_BenchmarkRet{messageRet}
		return true
	}
	
	
	response.Code = protocol.Code_InvalidRequest
	return true
}

