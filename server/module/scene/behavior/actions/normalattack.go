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
	b3 "github.com/magicsea/behavior3go"
	b3config "github.com/magicsea/behavior3go/config"
	b3core "github.com/magicsea/behavior3go/core"
)

//普通攻击
type NormalAttack struct {
	b3core.Action

	//targetID string
}

func (this *NormalAttack) Initialize(setting *b3config.BTNodeCfg) {
	this.Action.Initialize(setting)
}

func (this *NormalAttack) OnTick(tick *b3core.Tick) b3.Status {
	source := tick.GetTarget().(*mmo.Entity) //攻击者

	target := tick.Blackboard.Get("target", tick.GetTree().GetID(), "").(*mmo.Entity) //被攻击者
	//目标不存在执行失败
	if target == nil {
		return b3.FAILURE
	}

	mmo.Call(source.GetID(), "NormalAttack", target)
	return b3.SUCCESS
}

