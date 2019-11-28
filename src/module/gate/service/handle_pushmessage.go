// Code generated by aliensboot. DO NOT EDIT.
// source: gate_interface.proto
package service

import (
	"github.com/KylinHe/aliensboot-server/constant"
	"github.com/KylinHe/aliensboot-server/dispatch/rpc"
	"github.com/KylinHe/aliensboot-server/module/gate/cache"
	"github.com/KylinHe/aliensboot-server/module/gate/network"
	"github.com/KylinHe/aliensboot-server/module/gate/route"
	"github.com/KylinHe/aliensboot-server/module/gate/util"
	"github.com/KylinHe/aliensboot-server/protocol"
	"github.com/KylinHe/aliensboot-core/cluster/center"
	"github.com/KylinHe/aliensboot-core/protocol/base"
)

//
func handlePushMessage(authID int64, gateID string, request *protocol.PushMessage) {
	msgID := route.GetServiceSeq(request.GetService())
	if msgID == 0 {
		msgID = constant.SystemMsgId
	}
	msg := &base.Any{Id: msgID, Value: request.GetData()}
	if request.GetAll() {
		network.Manager.Broadcast(msg)
	} else if request.GetAuthId() != nil {
		// 允许转发 需要查找路由转发
		if request.GetRelay() {
			routes := make(util.String2Int64)
			for _, authID := range request.GetAuthId() {
				node := cache.GetAuthGateID(authID)
				if node != "" {
					routes.Put(node, authID)
				}
			}
			for node, authIDs := range routes {
				if node == center.ClusterCenter.GetNodeID() {
					network.Manager.PushMulti(authIDs, msg)
				} else {
					pushMessage := &protocol.PushMessage{
						AuthId:  authIDs,
						Data:    request.GetData(),
						Service: request.GetService(),
					}
					_ = rpc.Gate.PushMessage(node, pushMessage)
				}
			}
		} else {
			network.Manager.PushMulti(request.GetAuthId(), msg)
		}
	}
}
