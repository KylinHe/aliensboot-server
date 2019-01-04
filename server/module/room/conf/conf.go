package conf

import (
	"github.com/KylinHe/aliensboot-core/config"
)

var Config struct {
	Service  config.ServiceConfig //grpc
	Cache    config.CacheConfig   //redis
	Database config.DBConfig      //mongo
}
