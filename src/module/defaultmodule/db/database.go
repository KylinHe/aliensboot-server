package db

import (
	"github.com/KylinHe/aliensboot-core/database"
	"github.com/KylinHe/aliensboot-core/database/mongo"
	"github.com/KylinHe/aliensboot-core/exception"
	"github.com/KylinHe/aliensboot-core/log"
	"github.com/KylinHe/aliensboot-server/module/defaultmodule/conf"
	"github.com/KylinHe/aliensboot-server/protocol"
)

var _db = &mongo.Database{}

func Init() {
	err := _db.Init(conf.Config.Database)
	if err != nil {
		panic(err)
	}

	_db.SetErrorHandler(func (err error) {
		if err != mongo.ErrNotFound {
			log.Errorf("database err: %v", err)
			exception.GameException(protocol.Code_DBExcetpion)
		}
	})
	//EnsureTable("collection", &protocol.Collection{})
}

func Close() {
	//_db.Close()
}


func EnsureTable(name string, data database.IData) {
	err := _db.EnsureTable(name, data)
	if err != nil {
		log.Fatalf("ensure collection %v err : %v", name, err)
	}
}
