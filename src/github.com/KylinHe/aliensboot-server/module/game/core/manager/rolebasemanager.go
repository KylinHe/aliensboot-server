/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/7/26
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package manager

import (
	"github.com/KylinHe/aliensboot-server/protocol"
)

//角色资产管理
type RoleBaseManager struct {
	*protocol.Role
}

//初始化
func (this *RoleBaseManager) Init(role *protocol.Role) {
	this.Role = role
}

//更新数据库内存
func (this *RoleBaseManager) Update(role *protocol.Role) {
	role = this.Role
}

func (this *RoleBaseManager) ChangeNickname(nickname string) {
	this.Nickname = nickname
}
