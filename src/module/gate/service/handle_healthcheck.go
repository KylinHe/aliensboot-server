// Code generated by aliensboot. DO NOT EDIT.
// source: gate_interface.proto
package service

import (
	"github.com/KylinHe/aliensboot-server/module/gate/network"
	"github.com/KylinHe/aliensboot-server/protocol"
)

//
func handleHealthCheck(authID int64, gateID string, request *protocol.HealthCheck, response *protocol.HealthCheckRet) {
	response.Online = network.Manager.HealthCheck(request.GetAuthId())
}
