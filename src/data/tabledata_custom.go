/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2019/11/27
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package data

import "time"

// [账号白名单表]
type WhitelistAccount struct {
	Account string `json:"account"` //账号白名单

	Enable bool `json:"enable" form-type:"radio"` //是否启用

	Note string `json:"note"` //名称
}

// [ip白名单表]
type WhitelistIp struct {
	Ip string `json:"ip" form-type:"ip"` //ip地址

	Enable bool `json:"enable" form-type:"radio"` //是否启用

	Note string `json:"note"` // 描述信息
}

// [uid白名单表]
type WhitelistUid struct {
	Uid int64 `json:"uid" form-type:"number"` //uid

	Enable bool `json:"enable" form-type:"radio"` //是否启用

	Note string `json:"note"` // 描述信息
}

// [公告]
type NoticeData struct {
	Content string `json:"content"` // 内容

	PublicTime int64 `json:"publicTime"` // 发布时间
}

type MaintainBase struct {
	Open       bool      `json:"open"`      //维护状态
	StartTime  int64     `json:"startTime"` //维护的开始时间
	EndTime    int64     `json:"endTime"`   //维护的结束时间
	StartTime1 time.Time `json:"-"`         //服务器的开放时间戳
	EndTime1   time.Time `json:"-"`         //服务器的开放时间戳
}
