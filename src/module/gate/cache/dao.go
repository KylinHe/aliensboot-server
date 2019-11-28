/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/4/11
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package cache

import "github.com/KylinHe/aliensboot-core/common/util"

const (
	userGatePrefix string = "gate:" //
)

func getAuthGateKey(authID int64) string {
	return userGatePrefix + util.Int64ToString(authID)
}

//设置客户端所在的网关id
func SetAuthGateID(authID int64, gateID string) error {
	return _cache.SetData(getAuthGateKey(authID), gateID)
}

//清楚用户和网关的对应关系
func CleanAuthGateID(authID int64) error {
	return _cache.DelData(getAuthGateKey(authID))
}

//获取客户端所在的网关id
func GetAuthGateID(authID int64) string {
	result, _ := _cache.GetData(getAuthGateKey(authID))
	return result
}
