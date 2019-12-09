// Code generated by aliensboot. DO NOT EDIT.
// source: gate_interface.proto
package service

import (
	"github.com/KylinHe/aliensboot-server/module/gate/cache"
	"github.com/KylinHe/aliensboot-server/protocol"
	"github.com/KylinHe/aliensboot-core/cluster/center/service"
)

//
func handleGetAuthNode(ctx *service.Context, request *protocol.GetAuthNode, response *protocol.GetAuthNodeRet) {
	response.Node = cache.GetAuthGateID(request.GetAuthId())
	return
}
