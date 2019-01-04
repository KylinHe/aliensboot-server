package db

import (
	"github.com/KylinHe/aliensboot-core/database/mongo"
	"github.com/KylinHe/aliensboot-server/module/game/conf"
	"github.com/KylinHe/aliensboot-server/protocol"
)

var Database *mongo.Database = &mongo.Database{}

func Init() {
	err := Database.Init(conf.Config.Database)
	if err != nil {
		panic(err)
	}

	Database.EnsureTable("role", &protocol.Role{})

}

func Close() {
	Database.Close()
}
