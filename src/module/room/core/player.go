/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/11/13
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package core

import (
	"github.com/KylinHe/aliensboot-server/constant"
	"github.com/KylinHe/aliensboot-server/dispatch/rpc"
	"github.com/KylinHe/aliensboot-server/protocol"
	"github.com/gogo/protobuf/proto"
)

type Player struct {
	*protocol.Player

	gateID string //是否绑定了网关、玩家是否连接到服务器

	ready bool //玩家是否准备完毕
}

func (player *Player) SendMsg(data []byte) {
	pushMessage := &protocol.PushMessage{
		AuthID:  player.GetId(),
		Data:    data,
		Service: "room",
	}
	rpc.Gate.PushMessage(player.gateID, pushMessage)
}

func (player *Player) SendProtoMsg(message proto.Message) {
	sendData, _ := proto.Marshal(message)
	player.SendMsg(sendData)
}

func (player *Player) kick(kickType protocol.KickType) {
	rpc.Gate.KickOut("", &protocol.KickOut{
		AuthID:   player.GetId(),
		KickType: kickType,
	})
}

func (player *Player) Ready() {
	player.ready = true
}

func (player *Player) IsAnchor() bool {
	return player.HaveRole(constant.RoleAnchor)
}

func (player *Player) HaveRole(role int32) bool {
	return (player.GetRole() & role) != 0
}

func (player *Player) RemoveRole(role int32) {
	if player.HaveRole(role) {
		player.Role -= role
	}
}

func (player *Player) AddRole(role int32) {
	player.Role |= role
}

func (player *Player) IsReady() bool {
	return player.ready
}
