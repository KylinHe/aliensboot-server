/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2017/8/4
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package conf

import (
	"github.com/KylinHe/aliensboot-core/config"
)

//type Route struct {
//	Service string `json:"service"`
//	Seq     uint16 `json:"seq"`
//	Auth    bool   `json:"auth"`
//}

var Config struct {
	//Enable              bool   //网络模块是否开启
	Service          config.ServiceConfig
	Cache            config.CacheConfig
	TCP              config.TCPConfig
	WebSocket        config.WsConfig
	Http             config.HttpConfig
	SecretKey        string //
	AuthTimeout      float64
	HeartbeatTimeout float64
	Routes            map[string]uint16 //路由配置
}

//func Init(name string) {
//	Config.Service.Name = name
//}
