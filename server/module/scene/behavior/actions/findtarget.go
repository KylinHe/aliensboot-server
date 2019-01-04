/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/12/13
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package actions

import (
	"github.com/KylinHe/aliensboot-core/mmo"
	"github.com/KylinHe/aliensboot-server/module/scene/entity"
	b3 "github.com/magicsea/behavior3go"
	b3config "github.com/magicsea/behavior3go/config"
	b3core "github.com/magicsea/behavior3go/core"
)

//巡逻检查玩家
type FindTarget struct {
	b3core.Action
	target string
}

func (this *FindTarget) Initialize(setting *b3config.BTNodeCfg) {
	this.Action.Initialize(setting)
	this.target = setting.GetPropertyAsString("target")
}

func (this *FindTarget) OnTick(tick *b3core.Tick) b3.Status {
	source := tick.GetTarget().(*mmo.Entity)

	monster, ok := source.I.(*entity.Monster)
	if !ok {
		return b3.FAILURE
	}
	//targetType := tick.Blackboard.Get(this.targetType, tick.GetTree().GetID(), "").(mmo.EntityType)

	entity := monster.FindTarget(entity.TypePlayer)
	if entity == nil {
		return b3.FAILURE
	}

	tick.Blackboard.Set(this.target, entity, tick.GetTree().GetID(), "")
	return b3.SUCCESS
}
