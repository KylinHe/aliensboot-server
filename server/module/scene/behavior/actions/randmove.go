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
	"github.com/KylinHe/aliensboot-server/module/scene/entity"
	b3 "github.com/magicsea/behavior3go"
	b3config "github.com/magicsea/behavior3go/config"
	b3core "github.com/magicsea/behavior3go/core"
)

//随机移动
type RandMove struct {
	b3core.Action
}

func (this *RandMove) Initialize(setting *b3config.BTNodeCfg) {
	this.Action.Initialize(setting)
}

func (this *RandMove) OnTick(tick *b3core.Tick) b3.Status {
	f := tick.GetTarget().(*entity.Monster)
	f.RandMove()
	return b3.SUCCESS
}
