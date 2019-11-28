/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2019/9/23
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package conf

func IsMaintain() bool {
	return Maintain.Open
}

func EnsureWhitelistUid(uid int64) bool {
	return WhitelistUidData[uid]
}

// 测试账号
func EnsureWhitelistAccount(account string) bool {
	return WhitelistAccountData[account]
}

func EnsureWhitelistIp(ip string) bool {
	return WhitelistIpData[ip]
}

//func GetCheckState(uid int32, newUser bool, ip string) protocol.Code {
//	if Maintain.State == 1 {
//		if !clustercache.Cluster.IsUidWhiteList(uid) &&
//			!clustercache.Cluster.IsIpWhiteList(ip) {
//			return exception.SERVER_MAINTAIN
//		}
//	}
//
//	status := cache.UserCache.GetUserAttrInt32(uid, basecache.UPROP_STATUS)
//	//用户是否被封号
//	if byte(status) == db.USER_STATUS_NOT_AUTH {
//		return exception.USER_FORBIDDEN
//	}
//	return protocol.Code_Success
//}
