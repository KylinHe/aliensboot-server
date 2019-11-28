package core

import (
	"crypto/md5"
	"github.com/KylinHe/aliensboot-core/exception"
	"github.com/KylinHe/aliensboot-server/constant"
	"github.com/KylinHe/aliensboot-server/dispatch/lpc"
	"github.com/KylinHe/aliensboot-server/module/passport/conf"
	"github.com/KylinHe/aliensboot-server/module/passport/db"
	"github.com/KylinHe/aliensboot-server/protocol"
	"encoding/hex"
	"github.com/KylinHe/aliensboot-core/common/util"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func ValidateState(user *protocol.User, ip string) protocol.LoginResult {
	if user == nil {
		return protocol.LoginResult_invalidUser
	}
	if user.GetStatus() == constant.UserStateForbidden {
		return protocol.LoginResult_forbiddenUser
	}
	if user.Status == constant.UserStateBanned { // 封号
		exception.GameException(protocol.LoginResult_forbiddenUser)
	}
	if conf.IsMaintain() {
		// 白名单通过
		if conf.EnsureWhitelistIp(ip) || conf.EnsureWhitelistAccount(user.GetUsername()) || conf.EnsureWhitelistUid(user.GetId()) {
			return protocol.LoginResult_loginSuccess
		} else {
			return protocol.LoginResult_invalidMaintain
		}
	}
	return protocol.LoginResult_loginSuccess
}

func ResetUser(username string) {
	newUserName := username + "_" + bson.NewObjectId().Hex()
	passwordHash := PasswordHash(newUserName, conf.Config.DefaultChannelPassport)
	_ = db.UpdateUserUsername(username, newUserName, passwordHash)
}

func PasswordHash(username string, passwd string) string {
	//h.Write([]byte(passwd + userCache.Salt))
	return MD5Hash(username + MD5Hash(passwd))
}

func GetUserByUid(uid int64) *protocol.User {
	return db.QueryUser(uid)
}

func GetUserByUsername(username string) *protocol.User {
	return db.QueryUserByUsername(username)
}

func MD5Hash(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	md5Hash := hex.EncodeToString(h.Sum(nil))
	return md5Hash
}

func NewToken() string {
	return util.GenUUID()
}

/**
 *  新建用户
 */
func NewUser(username string, password string, channel string, channelUID string, openID string, avatar string) *protocol.User {
	passwordHash := PasswordHash(username, password)
	user := db.CreateUser(username, passwordHash, channel, channelUID, openID, avatar)
	lpc.LogHandler.AddLog(&protocol.LogRegister{
		Id:       bson.NewObjectId().Hex(),
		RoleId:   user.Id,
		Time:     time.Now().Unix(),
		Channel:  user.Channel,
		Platform: user.Channel,
	})
	return user
}

// 修改玩家状态
func ModifyUserStatus(uid int64, status int32) *protocol.User {
	return db.UpdateUserStatus(uid, status)
}
