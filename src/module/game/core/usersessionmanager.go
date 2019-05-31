/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/7/25
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package core

import (
	"github.com/KylinHe/aliensboot-core/exception"
	"github.com/KylinHe/aliensboot-server/protocol"
)

var UserManager = &userManager{users: make(map[int64]*UserSession)}

type userManager struct {
	users map[int64]*UserSession
}

//加载用户基础数据
func (this *userManager) EnsureUser(uid int64) *UserSession {
	session := this.users[uid]
	if session == nil {
		session = newUserSession(uid)
		this.users[uid] = session
	}
	return session
}

func (this *userManager) GetUser(uid int64) *UserSession {
	session := this.users[uid]
	if session == nil {
		exception.GameException(protocol.Code_ValidateException)
	}
	return session
}
