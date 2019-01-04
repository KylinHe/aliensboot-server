/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/11/14
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package task

import (
	"github.com/KylinHe/aliensboot-core/log"
	"github.com/KylinHe/aliensboot-core/module/base"
	"github.com/KylinHe/aliensboot-core/task"
	"github.com/KylinHe/aliensboot-server/module/hall/match"
)

func Init(skeleton *base.Skeleton) {
	//匹配检查
	cron, err := task.NewCronExpr("*/2 * * * * *")
	if err != nil {
		log.Error("init match timer error : %v", err)
	}
	skeleton.CronFunc(cron, match.Manager.TryMatch)
}
