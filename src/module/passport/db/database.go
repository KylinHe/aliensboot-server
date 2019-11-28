package db

import (
	"github.com/KylinHe/aliensboot-server/module/passport/conf"
	"github.com/KylinHe/aliensboot-server/protocol"
	"github.com/KylinHe/aliensboot-core/database"
	"github.com/KylinHe/aliensboot-core/database/mongo"
	"github.com/KylinHe/aliensboot-core/exception"
	"github.com/KylinHe/aliensboot-core/log"
)

var _db = &mongo.Database{}

func Init() {
	err := _db.Init(conf.Config.Database)
	if err != nil {
		panic(err)
	}
	_db.SetErrorHandler(func(err error) {
		if err != mongo.ErrNotFound {
			log.Errorf("passport database err: %v", err)
			exception.GameException(protocol.Code_DBExcetpion)
		}
	})
	EnsureTable("passport", &protocol.User{})
	//DatabaseHandler.Insert(&passport.User{Id:DatabaseHandler.GenTimestampId(&passport.User{}),Username:"hejialin",RegTime:time.Now()})
}

func Close() {
	_db.Close()
}

func EnsureTable(name string, data database.IData) {
	err := _db.EnsureTable(name, data)
	if err != nil {
		log.Fatalf("ensure collection %v err : %v", name, err)
	}
}
