package db

import (
	"github.com/KylinHe/aliensboot-server/module/oplog/conf"
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
			log.Errorf("database err: %v", err)
			exception.GameException(protocol.Code_DBExcetpion)
		}
	})
	EnsureTable("log_login", &protocol.LogLogin{})
	EnsureTable("log_logout", &protocol.LogLogout{})
	EnsureTable("log_register", &protocol.LogRegister{})
	EnsureTable("log_diamond", &protocol.LogDiamond{})
	EnsureTable("log_charge", &protocol.LogCharge{})

	RegisterDayTable(AnalysisTypeDau, &protocol.LogDayActiveUser{})
	RegisterDayTable(AnalysisTypeDv, &protocol.LogDayVisit{})
	RegisterDayTable(AnalysisTypeDr, &protocol.LogDayRegister{})
	RegisterDayTable(AnalysisTypeDrt, &protocol.LogDayRegisterTotal{})
	RegisterDayTable(AnalysisTypeDav, &protocol.LogDayAvgVisit{})
	RegisterDayTable(AnalysisTypeDat, &protocol.LogDayAvgTime{})

	RegisterDayTable(AnalysisTypeDrr1, &protocol.LogDayRegisterRetention1{})
	RegisterDayTable(AnalysisTypeDrr3, &protocol.LogDayRegisterRetention3{})
	RegisterDayTable(AnalysisTypeDrr7, &protocol.LogDayRegisterRetention7{})
	RegisterDayTable(AnalysisTypeDar1, &protocol.LogDayActiveRetention1{})
	RegisterDayTable(AnalysisTypeDar3, &protocol.LogDayActiveRetention3{})
	RegisterDayTable(AnalysisTypeDar7, &protocol.LogDayActiveRetention7{})
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
