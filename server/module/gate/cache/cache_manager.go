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
	"github.com/KylinHe/aliensboot-core/cache/redis"
	"github.com/KylinHe/aliensboot-core/config"
	"github.com/KylinHe/aliensboot-server/module/gate/conf"
)

var GateCache = &cacheManager{redisClient: &redis.RedisCacheClient{}}

type cacheManager struct {
	redisClient *redis.RedisCacheClient
}

func Init() {
	GateCache.Init(conf.Config.Cache)
}

func Close() {
	GateCache.Close()
}

func (this *cacheManager) Init(config config.CacheConfig) {
	this.redisClient = redis.NewRedisClient(config)
	this.redisClient.Start()
}

func (this *cacheManager) Close() {
	if this.redisClient != nil {
		this.redisClient.Close()
	}
}
