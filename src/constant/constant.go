/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/11/14
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package constant

const SystemAuthId int64 = -1
const SystemMsgId uint16 = 666

const (
	TestAppID = "0"

	RoleNone   int32 = 0
	RoleAnchor int32 = 1 << 0   //房主
	RolePlayer int32 = 1 << 1   //玩家
	RoleViewer int32 = 1 << 2   //ob

	RoleAll = RoleAnchor | RolePlayer | RoleViewer | RoleNone

	AnySeat int32 = -1 //空闲座位

	//0开放座位, 1锁定座位 2离开座位 3替换座位
	OptUnlockSeat int32 = 0
	OptLockSeat   int32 = 1
	OptLeaveSeat  int32 = 2
	OptChangeSeat int32 = 3 //替换座位
)

const (
	ModulePassportId uint16 = 1
	HeaderIP         string = "ip"
)

const (
	UserStateNone      int = 0
	UserStateForbidden     = 1
	UserStateBanned        = 2
)