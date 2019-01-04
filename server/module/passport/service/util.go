/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/4/2
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package service

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/KylinHe/aliensboot-core/common/util"
)

func MD5Hash(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	md5Hash := hex.EncodeToString(h.Sum(nil))
	return md5Hash
}

func PasswordHash(username string, passwd string) string {
	//h.Write([]byte(passwd + userCache.Salt))
	return MD5Hash(username + MD5Hash(passwd))
}

func NewToken() string {
	return util.GenUUID()
}
