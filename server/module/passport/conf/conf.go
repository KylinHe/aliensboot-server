package conf

import (
	"github.com/KylinHe/aliensboot-core/config"
	"time"
)

var Config struct {
	Service  config.ServiceConfig
	Cache    config.CacheConfig
	Database config.DBConfig

	DefaultChannelPassport string
	TokenExpireTime        int64
	//HTTPAddress       string
	AppKey string
}

func Init() {
	if Config.TokenExpireTime <= 0 {
		//默认过期时间七天
		Config.TokenExpireTime = int64(7 * 24 * time.Hour)
	}
}

//func GetTokenExpireTimestamp() int64 {
//	return time.Now().Add(time.Duration(Config.TokenExpireTime)).Unix()
//}
