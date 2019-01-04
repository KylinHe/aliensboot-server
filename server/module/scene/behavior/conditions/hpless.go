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
	"github.com/KylinHe/aliensboot-core/mmo"
	b3 "github.com/magicsea/behavior3go"
	b3config "github.com/magicsea/behavior3go/config"
	b3core "github.com/magicsea/behavior3go/core"
)

//HpLess
type HpLess struct {
	b3core.Condition

	rate float32 //血量低于百分比
}

func (this *HpLess) Initialize(setting *b3config.BTNodeCfg) {
	this.Condition.Initialize(setting)
	this.rate = float32(setting.GetProperty("rate"))
}

func (this *HpLess) OnTick(tick *b3core.Tick) b3.Status {
	f := tick.GetTarget().(*mmo.Entity)

	maxhp := f.GetFloat32("maxhp")
	//没有血量属性 返回失败
	if maxhp == 0 {
		return b3.FAILURE
	}

	hp := f.GetFloat32("hp")

	rate := hp / maxhp

	if rate < this.rate {
		return b3.SUCCESS
	}
	return b3.FAILURE
}
