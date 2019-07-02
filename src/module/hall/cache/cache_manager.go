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
	"github.com/KylinHe/aliensboot-core/exception"
	"github.com/KylinHe/aliensboot-core/log"
	"github.com/KylinHe/aliensboot-server/module/hall/conf"
	"github.com/KylinHe/aliensboot-server/protocol"
)

var Manager = &cacheManager{redisClient: &redis.RedisCacheClient{}}

type cacheManager struct {
	redisClient *redis.RedisCacheClient
}

func Init() {
	Manager.Init(conf.Config.Cache)
	////TODO 加载数据库到缓存
	//if PassportCache.SetNX(FLAG_LOADUSER, 1) {
	//	var users []*protocol.User
	//	db.DatabaseHandler.QueryAll(&protocol.User{}, &users)
	//	for _, user := range users {
	//		PassportCache.SetUsernameUidMapping(user.GetUsername(), user.GetId())
	//		PassportCache.HSetUser(user.GetId(), user)
	//	}
	//}
}

func Close() {
	Manager.Close()
}

func (this *cacheManager) Init(config config.CacheConfig) {
	this.redisClient = redis.NewRedisClient(config)
	this.redisClient.SetErrorHandler(func (err error, command string, args... interface{}) {
		if err != redis.ErrNil {
			log.Errorf("[%v] %v err: %v", command, args, err)
			exception.GameException(protocol.Code_DBExcetpion)
		}
	})
	this.redisClient.Start()
}

func (this *cacheManager) Close() {
	if this.redisClient != nil {
		this.redisClient.Close()
	}
}

func (this *cacheManager) SetNX(key string, value interface{}) bool {
	result, _ := this.redisClient.SetNX(key, value)
	return result
}
