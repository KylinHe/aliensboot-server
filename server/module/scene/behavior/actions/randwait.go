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
	b3 "github.com/magicsea/behavior3go"
	b3config "github.com/magicsea/behavior3go/config"
	b3core "github.com/magicsea/behavior3go/core"
	"math/rand"
	"time"
)

//等待随机时间
type RandWait struct {
	b3core.Action
	minTime int64 //最少等待时间
	maxTime int64 //最长等待时间
}

func (this *RandWait) Initialize(setting *b3config.BTNodeCfg) {
	this.Action.Initialize(setting)
	this.minTime = setting.GetPropertyAsInt64("waitmin")
	this.maxTime = setting.GetPropertyAsInt64("waitmax")
}

func (this *RandWait) OnOpen(tick *b3core.Tick) {
	var startTime int64 = time.Now().UnixNano() / 1000000
	tick.Blackboard.Set("startTime", startTime, tick.GetTree().GetID(), this.GetID())
	end := this.minTime + rand.Int63n(this.maxTime-this.minTime)
	tick.Blackboard.Set("endTime", startTime+end, tick.GetTree().GetID(), this.GetID())
}

func (this *RandWait) OnTick(tick *b3core.Tick) b3.Status {
	var currTime int64 = time.Now().UnixNano() / 1000000
	var endTime = tick.Blackboard.GetInt64("endTime", tick.GetTree().GetID(), this.GetID())

	if currTime > endTime {
		return b3.SUCCESS
	}
	return b3.RUNNING
}
