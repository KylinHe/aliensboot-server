package conf

import (
	"github.com/KylinHe/aliensboot-core/config"
)

var Config struct {
	Service  config.ServiceConfig
	Cache    config.CacheConfig
	Database config.DBConfig
}
