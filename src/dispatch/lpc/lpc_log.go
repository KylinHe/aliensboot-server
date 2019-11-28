/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/5/10
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package lpc

import (
	"github.com/KylinHe/aliensboot-core/database"
	"github.com/KylinHe/aliensboot-server/constant"
	"github.com/KylinHe/aliensboot-server/module/oplog"
)

var LogHandler = &logHandler{}

type logHandler struct {
}

func (handler *logHandler) AddLog(data database.IData) {
	oplog.ChanRPC.Go(constant.LogCommand, data)
}
