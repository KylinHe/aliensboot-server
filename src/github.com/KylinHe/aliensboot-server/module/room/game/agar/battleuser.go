/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2019/2/15
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package agar

import "github.com/KylinHe/aliensboot-server/module/scene/entity"

/**
 * 玩家战斗对象
 */
type BattleUser struct {
	*entity.Player

	color int32 //皮肤id

	ballCount int32 //当前球的数量

	balls map[int32]*Ball //拥有的球

	battle *AgarGame

}
