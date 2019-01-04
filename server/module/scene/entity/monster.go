/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/12/4
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package entity

import (
	"github.com/KylinHe/aliensboot-core/log"
	"github.com/KylinHe/aliensboot-core/mmo"
	"github.com/KylinHe/aliensboot-core/mmo/core"
	"github.com/KylinHe/aliensboot-server/module/scene/constant"
)

const (
	TypeMonster mmo.EntityType = "Monster"
)


//
type Monster struct {
	mmo.Entity   // Entity type should always inherit entity.Entity

	target *mmo.Entity  //当前锁定的目标

}

func (monster *Monster) DescribeEntityType(desc *core.EntityDesc) {
	//monster.SetAI()

	desc.SetUseAOI(true, 100)

	desc.DefineAttr(constant.AttrLevel, core.AttrAllClient| core.AttrPersist)
	desc.DefineAttr(constant.AttrHp, core.AttrAllClient| core.AttrPersist)
	desc.DefineAttr(constant.AttrMaxHp, core.AttrAllClient| core.AttrPersist)
	desc.DefineAttr(constant.AttrAction, core.AttrAllClient)
}

func (monster *Monster) TestCall(param1 string, param2 int32, param3 []string) {
	log.Debugf("test call : %v %v %v", param1, param2, param3)
}

//随机移动
func (monster *Monster) RandMove() {
	log.Debugf("test call : %v %v %v", )
}

func (monster *Monster) Attack() {
	log.Debugf("test call : %v %v %v", )
}

//检查范围内是否有目标存在
func (monster *Monster) FindTarget(targetType mmo.EntityType) *mmo.Entity {
	interest := monster.GetInterest()
	for target, _ := range interest {
		if target.GetType() == targetType {
			return target
		}
	}
	return nil
}

