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

//var PassportCache = &cacheManager{RedisCacheClient: &redis.RedisCacheClient{}}
//
//type cacheManager struct {
//	*redis.RedisCacheClient
//}
//
//func Init() {
//	PassportCache.Init(conf.Config.Cache)
//
//	count := 0
//
//	// 清空缓存的情况下需要重新加载缓存
//	if ok, _ := PassportCache.SetNX(FlagLoadUser, 1); ok {
//		log.Debug("load passport data to redis cache start...")
//		var users []*protocol.User
//		var commands []redis.Command
//		_ = db.Database.QueryAllLimit(&protocol.User{}, &users, 10000, func(data interface{}) bool {
//			currLen := len(users)
//			if currLen == 0 {
//				return true
//			}
//			commands = make([]redis.Command, currLen)
//			start := time.Now()
//			for index, user := range users {
//				commands[index] = redis.Command{Args: []interface{}{redis.OP_SET, PassportCache.GetUsernamePrefix(user.GetUsername()), user.GetId()}}
//				//key, hash, err := PassportCache.OGet(user, user.GetId())
//				//if err != nil {
//				//	log.Error(err)
//				//	continue
//				//}
//				//commands[currLen + index] = redis.Command{Args:[]interface{}{redis.OP_H_MSET, key, hash}}
//				//_ = PassportCache.SetUsernameUidMapping(user.GetUsername(), user.GetId())
//				//_ = PassportCache.HSetUser(user.GetId(), user)
//			}
//
//			// 使用pipeline优化批量处理
//			_ = PassportCache.PipelineCommands(commands)
//			count += currLen
//			log.Debugf("load passport data to redis cache ... %v - time %v", count, time.Now().Sub(start).Seconds())
//			return false
//		})
//		log.Debug("load passport data to redis cache done...")
//	}
//}
//
//func Close() {
//	_ = PassportCache.Close()
//}
//
//func Init(config config.CacheConfig) {
//	this.RedisCacheClient = redis.NewRedisClient(config)
//	this.RedisCacheClient.SetErrorHandler(func(err error, command string, args ...interface{}) {
//		if err != redis.ErrNil {
//			log.Errorf("[%v] %v err: %v", command, args, err)
//			exception.GameException(protocol.Code_DBExcetpion)
//		}
//	})
//	this.RedisCacheClient.Start()
//}
