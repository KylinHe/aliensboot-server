package db

import (
	"github.com/KylinHe/aliensboot-core/database/mongo"
	"github.com/KylinHe/aliensboot-core/log"
	"github.com/KylinHe/aliensboot-server/module/game/conf"
	"github.com/KylinHe/aliensboot-server/protocol"
)

var Database *mongo.Database = &mongo.Database{}

func Init() {
	err := Database.Init(conf.Config.Database)
	if err != nil {
		panic(err)
	}

	EnsureTable("role", &protocol.Role{})

}

func Close() {
	Database.Close()
}


func EnsureTable(name string, data interface{}) {
	err := Database.EnsureTable(name, data)
	if err != nil {
		log.Fatalf("ensure collection %v err : %v", name, err)
	}
}
