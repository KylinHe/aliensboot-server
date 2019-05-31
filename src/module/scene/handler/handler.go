/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/12/4
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package handler

import (
	"errors"
	"fmt"
	"github.com/KylinHe/aliensboot-core/common/util"
	"github.com/KylinHe/aliensboot-core/mmo"
	"github.com/KylinHe/aliensboot-core/module/base"
	"github.com/KylinHe/aliensboot-server/dispatch/rpc"
	"github.com/KylinHe/aliensboot-server/module/scene/cache"
	"github.com/KylinHe/aliensboot-server/module/scene/conf"
	"github.com/KylinHe/aliensboot-server/module/scene/entity"
	"github.com/KylinHe/aliensboot-server/protocol"
)

type AliensEntityHandler struct {
	timerManager *util.TimerManager
}

//持久化
func (*AliensEntityHandler) Save(entityID mmo.EntityID, entityType mmo.EntityType, data map[string]interface{}, callback func()) error {
	//lpc.DBServiceProxy.ForceUpdate(&db.Entity{ID:entityID, Type:entityType, Data:data}, db.Database)
	return nil
}

func (*AliensEntityHandler) Load(entityID mmo.EntityID) (map[string]interface{}, error) {
	//data := &db.Entity{ID:entityID}
	//db.Database.QueryOne()
	return nil, nil
}

func (*AliensEntityHandler) CallRemote(entityID mmo.EntityID, method string, args [][]byte) error {
	node, err := cache.Manager.GetEntityNode(string(entityID))
	if err != nil {
		return err
	}
	if node == "" {
		return errors.New(fmt.Sprintf("call remote entity err, entity %v is not found", entityID))
	}
	rpc.Scene.EntityCall(node, &protocol.EntityCall{
		EntityID: string(entityID),
		Method:   method,
		Args:     args,
	})
	return nil
}

func (*AliensEntityHandler) MigrateRemote(spaceID mmo.EntityID, entityID mmo.EntityID, data []byte) error {
	space := string(spaceID)
	node, err := cache.Manager.GetSpaceNode(space)
	if err != nil {
		return err
	}
	if node == "" {
		return errors.New(fmt.Sprintf("migrate err, space %v is not found", space))
	}

	rpc.Scene.MigrateIn(node, &protocol.MigrateIn{
		SpaceID:  space,
		EntityID: string(entityID),
		Data:     data,
	})
	return nil
}

func (handler *AliensEntityHandler) GetTimerManager() *util.TimerManager {
	return handler.timerManager
}

func InitSpace() {
	space, _ := mmo.CreateSpace(entity.TypeGameSpace, mmo.EntityID(conf.Config.Space))
	cache.Manager.SetSpaceNode(string(space.GetID()))

	//space2, _ := mmo.CreateSpace(entity.TypeGameSpace, "space2")
	//cache.Manager.SetSpaceNode(string(space2.GetID()))
}

//初始化space
func Init(skeleton *base.Skeleton) {

	timerManager := util.NewTimerManager()
	skeleton.SetTick(timerManager.Tick)

	mmo.RegisterEntityHandler(&AliensEntityHandler{timerManager: timerManager})
	mmo.RegisterSpace(&entity.GameSpace{})
	mmo.RegisterEntity(&entity.Monster{})
	mmo.RegisterEntity(&entity.Player{})

	InitSpace()

	//entity := mmo.CreateEntity("Monster", space, unit.Vector{0,0,0})
	//log.Debugf("entity %v", entity.GetID())

	//mmo.Call(entity.GetID(), "TestCall", "1", 1, []string{"3"})

}
