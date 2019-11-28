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

//import (
//	"github.com/KylinHe/aliensboot-server/protocol"
//)
//
//const (
//	FlagLoadUser      string = "flag:user_load" //标识，是否加载用户数据到缓存
//	UserNameKeyPrefix string = "username:"
//)
//
//var user = &protocol.User{}
//
////设置用户会话token
//func SetUserToken(uid int64, token string) error {
//	return this.OSetFieldByID(user, uid, "token", token)
//}
//
////获取用户会话token
//func GetUserToken(uid int64) (string, error) {
//	return this.OGetFieldByID(user, uid, "token")
//}
//
//////设置用户头像
////func SetUserAvatar(uid int64, avatar string) error {
////	return this.OSetFieldByID(user, uid, "Avatar", avatar)
////}
////
//////获取用户头像
////func GetUserAvatar(uid int64) (string, error) {
////	return this.OGetFieldByID(user, uid, "Avatar")
////}
//
//func GetUserState(uid int64) (string, error) {
//	return this.OGetFieldByID(user, uid, "status")
//}
//
////用户名是否存在
//func IsUsernameExist(username string) (bool, error) {
//	uid, err := this.GetUidByUsername(username)
//	return uid != 0, err
//}
//
//func SetUsernameUidMapping(username string, uid int64) error {
//	return this.SetData(this.GetUsernamePrefix(username), uid)
//}
//
//func RemoveUsernameUidMapping(username string) error {
//	return this.DelData(this.GetUsernamePrefix(username))
//}
//
//func GetUidByUsername(username string) (int64, error) {
//	return this.GetDataInt64(this.GetUsernamePrefix(username))
//}
//
//func GetUsernamePrefix(username string) string {
//	return UserNameKeyPrefix + username
//}
//
////获取用户所有信息数据
//func HSetUser(uid int64, data interface{}) error {
//	return this.OSetByID(data, uid)
//}
//
////设置用户所有信息数据
//func HGetUser(uid int64, data interface{}) error {
//	return this.OGetByID(data, uid)
//}
//
////用户是否存在
//func IsUserExist(uid int64) bool {
//	result, _ := this.OExists(user, uid)
//	return result
//}
