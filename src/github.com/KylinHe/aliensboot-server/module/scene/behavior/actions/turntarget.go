/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/12/14
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

//移向目标
type TurnTarget struct {
	b3core.Action
	target string
}

func (this *TurnTarget) Initialize(setting *b3config.BTNodeCfg) {
	this.Action.Initialize(setting)
	this.target = setting.GetPropertyAsString("target")
}

func (this *TurnTarget) OnTick(tick *b3core.Tick) b3.Status {
	target := tick.Blackboard.Get(this.target, "", "").(*mmo.Entity)
	if target == nil {
		return b3.FAILURE
	}

	f := tick.GetTarget().(*mmo.Entity)
	if !f.IsInterestedIn(target) {
		tick.Blackboard.Set(this.target, nil, "", "")
		return b3.FAILURE
	}

	//转向目标
	f.FaceTo(target)
	return b3.SUCCESS
}
