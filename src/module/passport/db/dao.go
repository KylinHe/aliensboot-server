/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/4/12
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package db

import (
	"github.com/KylinHe/aliensboot-server/protocol"
	"github.com/KylinHe/aliensboot-core/database/mongo"
	"time"
)

func CreateUser(username string, password string, channel string, channelUid string, openId string, avatar string) *protocol.User {
	user := &protocol.User{
		Username:   username,
		Password:   password,
		Salt:       "",
		Channel:    channel,
		Channeluid: channelUid,
		Mobile:     "",
		Openid:     openId,
		Status:     0,
		Avatar:     avatar,
		RegTime:    time.Now().Unix(),
	}
	_ = _db.Insert(user)
	return user
}

func QueryUserByUsername(username string) *protocol.User {
	user := &protocol.User{}
	err := _db.QueryOneCondition(user, "username", username)
	if err == mongo.ErrNotFound {
		return nil
	}
	return user
}

func QueryUser(uid int64) *protocol.User {
	user := &protocol.User{}
	err := _db.QueryOne(user)
	if err == mongo.ErrNotFound {
		return nil
	}
	return user
}

func UpdateUserStatus(uid int64, status int32) *protocol.User {
	// TODO 改成更新语句
	user := QueryUser(uid)
	if user == nil {
		return user
	}
	user.Status = status
	_ = _db.UpdateOne(user)
	return user
}

func UpdateUserStatusByUsername(username string, status int32) error {
	// TODO 改成更新语句
	user := QueryUserByUsername(username)
	if user == nil {
		return mongo.ErrNotFound
	}
	user.Status = status
	return _db.UpdateOne(user)
}

func UpdateUserUsername(username string, newUsername string, newPassword string) error {
	// TODO 改成更新语句
	user := QueryUserByUsername(username)
	if user == nil {
		return mongo.ErrNotFound
	}
	user.Username = newUsername
	user.Password = newPassword
	return _db.UpdateOne(user)

}
