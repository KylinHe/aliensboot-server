// Code generated by aliensboot. DO NOT EDIT.
// source: gate_interface.proto
package service

import (
	"github.com/KylinHe/aliensboot-server/module/gate/network"
	"github.com/KylinHe/aliensboot-server/protocol"
	"github.com/KylinHe/aliensboot-core/cluster/center/service"
)

//
func handleHealthCheck(ctx *service.Context, request *protocol.HealthCheck, response *protocol.HealthCheckRet) {
	response.Online = network.Manager.HealthCheck(request.GetAuthId())
}
