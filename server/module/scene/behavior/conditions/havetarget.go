/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/12/13
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package conditions

import (
	b3 "github.com/magicsea/behavior3go"
	b3config "github.com/magicsea/behavior3go/config"
	b3core "github.com/magicsea/behavior3go/core"
)

//是否发现目标
type HaveTarget struct {
	b3core.Condition
	targetID string
}

func (this *HaveTarget) Initialize(setting *b3config.BTNodeCfg) {
	this.Condition.Initialize(setting)
	this.targetID = setting.GetPropertyAsString("targetID")
}

func (this *HaveTarget) OnTick(tick *b3core.Tick) b3.Status {
	targetID := tick.Blackboard.Get(this.targetID, tick.GetTree().GetID(), "")
	if targetID == "" {
		return b3.FAILURE
	} else {
		return b3.SUCCESS
	}
}
