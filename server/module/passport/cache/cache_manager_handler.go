/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/4/12
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package cache

import (
	"github.com/KylinHe/aliensboot-server/protocol"
)

const (
	FLAG_LOADUSER       string = "test_flag:user_load" //标识，是否加载用户数据到缓存
	USERNAME_KEY_PREFIX string = "test_username:"
)

var user = &protocol.User{}

//设置用户会话token
func (this *cacheManager) SetUserToken(uid int64, token string) error {
	return this.redisClient.OSetFieldByID(user, uid, "Token", token)
}

//获取用户会话token
func (this *cacheManager) GetUserToken(uid int64) (string, error) {
	return this.redisClient.OGetFieldByID(user, uid, "Token")
}

//设置用户头像
func (this *cacheManager) SetUserAvatar(uid int64, avatar string) error {
	return this.redisClient.OSetFieldByID(user, uid, "Avatar", avatar)
}

//获取用户头像
func (this *cacheManager) GetUserAvatar(uid int64) (string, error) {
	return this.redisClient.OGetFieldByID(user, uid, "Avatar")
}

//设置用户会话token
func (this *cacheManager) SetUserNickname(uid int64, nickname string) error {
	return this.redisClient.OSetFieldByID(user, uid, "Nickname", nickname)
}

//获取用户昵称
func (this *cacheManager) GetUserNickname(uid int64) (string, error) {
	return this.redisClient.OGetFieldByID(user, uid, "Nickname")
}

//用户是否在线
func (this *cacheManager) GetUserOnline(uid int64) (bool, error) {
	return this.redisClient.OGetBoolFieldByID(user, uid, "Online")
}

//用户名是否存在
func (this *cacheManager) IsUsernameExist(username string) (bool, error) {
	uid, err := this.GetUidByUsername(username)
	return uid != 0, err
}

func (this *cacheManager) SetUsernameUidMapping(username string, uid int64) error {
	return this.redisClient.SetData(USERNAME_KEY_PREFIX+username, uid)
}

func (this *cacheManager) GetUidByUsername(username string) (int64, error) {
	return this.redisClient.GetDataInt64(USERNAME_KEY_PREFIX + username)
}

//获取用户所有信息数据
func (this *cacheManager) HSetUser(uid int64, data interface{}) {
	this.redisClient.OSetByID(data, uid)
}

//设置用户所有信息数据
func (this *cacheManager) HGetUser(uid int64, data interface{}) {
	this.redisClient.OGetByID(data, uid)
}

//用户是否存在
func (this *cacheManager) IsUserExist(uid int64) bool {
	result, _ := this.redisClient.OExists(user, uid)
	return result
}
