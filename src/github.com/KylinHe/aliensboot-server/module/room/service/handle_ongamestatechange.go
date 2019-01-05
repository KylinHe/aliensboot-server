// Code generated by aliensbot. DO NOT EDIT.
// source: room_interface.proto
package service

import (
	"github.com/KylinHe/aliensboot-server/module/room/core"
	"github.com/KylinHe/aliensboot-server/protocol"
)

//
func handleOnGameStateChange(authID int64, gateID string, request *protocol.OnGameStateChange, response *protocol.OnGameStateChangeRet) {
	response.Code = core.RoomManager.ChangeGameState(authID, request.GetState())

}