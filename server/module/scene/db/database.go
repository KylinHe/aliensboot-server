package db

import (
	"github.com/KylinHe/aliensboot-core/database/mongo"
	"github.com/KylinHe/aliensboot-server/module/scene/conf"
)

var Database *mongo.Database = &mongo.Database{}

func Init() {
	err := Database.Init(conf.Config.Database)
	if err != nil {
		panic(err)
	}
	Database.EnsureTable("entity", &Entity{})
}

func Close() {
	Database.Close()
}
