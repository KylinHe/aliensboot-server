/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2017/3/29
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package cache

import (
	"github.com/KylinHe/aliensboot-server/module/oplog/conf"
	"github.com/KylinHe/aliensboot-server/protocol"
	"github.com/KylinHe/aliensboot-core/cache/redis"
	"github.com/KylinHe/aliensboot-core/exception"
	"github.com/KylinHe/aliensboot-core/log"
)

var _cache = &redis.RedisCacheClient{}

func Init() {
	_cache = redis.NewRedisClient(conf.Config.Cache)
	_cache.SetErrorHandler(func(err error, command string, args ...interface{}) {
		if err != redis.ErrNil {
			log.Errorf("[%v] %v err: %v", command, args, err)
			exception.GameException(protocol.Code_DBExcetpion)
		}
	})
	_cache.Start()
}

func Close() {
	_ = _cache.Close()
}
