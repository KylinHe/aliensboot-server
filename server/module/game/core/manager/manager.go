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
	"github.com/KylinHe/aliensboot-core/common/util"
	"github.com/KylinHe/aliensboot-core/exception"
	"github.com/KylinHe/aliensboot-core/log"
	"github.com/KylinHe/aliensboot-server/module/game/db"
	"github.com/KylinHe/aliensboot-server/protocol"
	"reflect"
)

func NewRoleManager(uid int64) *RoleManager {
	var role = &protocol.Role{}
	err := db.Database.QueryOneCondition(role, "uid", uid)
	if err != nil {
		//创建数据
		role.Uid = uid
		role.Nickname = "蛇皮" + util.Int64ToString(uid)
		err1 := db.Database.Insert(role)
		if err1 != nil {
			log.Errorf("insert role error : %v", err1)
			exception.GameException(protocol.Code_DBExcetpion)
		}
	}
	dataManager := &RoleManager{data: role}
	dataManager.Init()
	return dataManager
}

//角色数据管理
type RoleManager struct {
	data *protocol.Role

	RoleBaseManager
}

func (this *RoleManager) IsRole(roleID int64) bool {
	return this.data.Id == roleID
}

func (this *RoleManager) GetData() *protocol.Role {
	return this.data
}

//初始化
func (this *RoleManager) Init() {
	mutable := reflect.ValueOf(this).Elem()
	params := make([]reflect.Value, 1)
	//数据管理类操作副本数据，这样更新的时候能够做比对增量更新
	params[0] = reflect.ValueOf(this.data)
	for i := 0; i < mutable.NumField(); i++ {
		f := mutable.Field(i)
		initMethod := f.Addr().MethodByName("Init")
		if initMethod.IsValid() {
			initMethod.Call(params)
		}
	}
}

//更新本地缓存
func (this *RoleManager) Update() {
	mutable := reflect.ValueOf(this).Elem()
	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(this.data)
	for i := 0; i < mutable.NumField(); i++ {
		f := mutable.Field(i)
		initMethod := f.Addr().MethodByName("Update")
		if initMethod.IsValid() {
			initMethod.Call(params)
		}
	}
}
