/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/12/4
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package entity

import (
	"github.com/KylinHe/aliensboot-core/common/util"
	"github.com/KylinHe/aliensboot-core/log"
	"github.com/KylinHe/aliensboot-core/mmo"
	"github.com/KylinHe/aliensboot-core/mmo/core"
	"github.com/KylinHe/aliensboot-core/mmo/unit"
	"github.com/KylinHe/aliensboot-server/dispatch/rpc"
	"github.com/KylinHe/aliensboot-server/module/scene/conf"
	"github.com/KylinHe/aliensboot-server/module/scene/constant"
	"github.com/KylinHe/aliensboot-server/module/scene/utils"
	"github.com/KylinHe/aliensboot-server/protocol"
	"time"
)

const (
	TypePlayer mmo.EntityType = "Player"

)

func GetPlayerID(authID int64) mmo.EntityID {
	return mmo.EntityID("P_" + util.Int64ToString(authID))
}

//
type Player struct {

	mmo.Entity   // Entity type should always inherit entity.Entity

	syncTimerID mmo.EntityTimerID
	releaseTimerID mmo.EntityTimerID

}

func (player *Player) DescribeEntityType(desc *core.EntityDesc) {
	//视野范围
	desc.SetUseAOI(true, 500)

	desc.DefineAttr(constant.AttrUid, core.AttrAllClient| core.AttrPersist) //用户id
	desc.DefineAttr(constant.AttrGateid, core.AttrClient)	//网关id
	desc.DefineAttr(constant.AttrLevel, core.AttrAllClient| core.AttrPersist)
	desc.DefineAttr(constant.AttrHp, core.AttrAllClient| core.AttrPersist)
	desc.DefineAttr(constant.AttrMaxHp, core.AttrAllClient| core.AttrPersist)
	desc.DefineAttr(constant.AttrAction, core.AttrAllClient)

}


func (player *Player) OnMigrateOut() {
	log.Debugf("handler migrateOut %v: %v - %v", player.GetSpaceID(), player.GetUid(), player.GetGateID())
}

func (player *Player) OnMigrateIn() {
	log.Debugf("handler migrateIn %v: %v - %v", player.GetSpaceID(), player.GetUid(), player.GetGateID())
	player.onLoad()
	player.syncTimerID = player.AddTimer(200 * time.Millisecond, true, "SyncData")
}


func (player *Player) Login(authID int64, gateID string) {
	log.Debugf("handler login %v - %v", authID, gateID)
	player.Set(constant.AttrUid, authID)
	player.Set(constant.AttrGateid, gateID)

	player.onLoad()

	//取消定时释放
	player.CancelTimer(player.releaseTimerID)
	//玩家每100ms同步一次数据
	player.syncTimerID = player.AddTimer(200 * time.Millisecond, true, "SyncData")
}

func (player *Player) Logout() {
	player.Set(constant.AttrGateid, "")
	player.CancelTimer(player.syncTimerID)

	//开启1分钟释放
	player.releaseTimerID = player.AddTimer(10 * time.Second, false, "Release")
}

func (player *Player) onLoad() {
	gateID := player.GetGateID()
	authID := player.GetUid()
	syncMessage := &protocol.Response{
		Scene:&protocol.Response_ScenePush {
			ScenePush:&protocol.ScenePush{
				SpaceID:string(player.GetSpaceID()),
				Entity:utils.BuildEntity(player.Entity, true),
			},
		},
	}

	//玩家的消息绑定到当前服务器节点
	rpc.Gate.BindService1(gateID, authID, conf.GetServiceName())
	rpc.Gate.Push(conf.GetServiceName(), authID, gateID, syncMessage)

}

func (player *Player) Move_Client(x string, y string) {
	player.SetPosition(unit.Vector{X:unit.Coord(util.StringToFloat32(x)), Y:unit.Coord(util.StringToFloat32(y)), Z:0})
}

//释放玩家内存
func (player *Player) Release() {
	player.Destroy()
}

////sync self 发送自己的玩家数据
func (player *Player) SyncData() {
	if !player.IsOnline() {
		return
	}

	interest := player.GetInterest()
	var entities = make([]*protocol.Entity, len(interest))

	index := 0
	for entity, _ := range interest {
		entities[index] = utils.BuildEntity(*entity, entity.GetID() == player.GetID())
		index ++
	}

	syncMessage := &protocol.Response{
		Scene:&protocol.Response_EntityPush{
			EntityPush:&protocol.EntityPush{
				Neighbors:entities,
			},
		},
	}

	rpc.Gate.Push(conf.GetServiceName(), player.GetUid(), player.GetGateID(), syncMessage)
}

func (player *Player) IsOnline() bool {
	return player.GetUid() > 0 && player.GetGateID() != ""
}

func (player *Player) GetUid() int64 {
	return player.GetInt64(constant.AttrUid)
}

func (player *Player) GetGateID() string {
	return player.GetString(constant.AttrGateid)
}
