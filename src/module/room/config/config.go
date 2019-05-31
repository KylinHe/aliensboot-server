/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/11/13
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package config

type RoomConfig struct {
	AppID    string //游戏类型id
	MaxSeat  int    //最大的座位数量
	Viewer   bool   //是否允许观众模式
	Anchor   bool   //支持房主管理
}
