// Code generated by aliensboot. DO NOT EDIT.
// source: room_interface.proto
package service

import (
	"github.com/KylinHe/aliensboot-server/module/room/core"
	"github.com/KylinHe/aliensboot-server/protocol"
)




//
func handleRoomClose(authID int64, gateID string, request *protocol.RoomClose) {
	core.RoomManager.CloseRoom(authID, request.GetRoomID())
}
