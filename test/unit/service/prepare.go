package service

import (
	"github.com/KylinHe/aliensboot-core/aliensboot"
	"github.com/KylinHe/aliensboot-core/module/database"
	"github.com/KylinHe/aliensboot-server/module/defaultmodule"
	"github.com/KylinHe/aliensboot-server/module/gate"
	"github.com/KylinHe/aliensboot-server/module/oplog"
	"github.com/KylinHe/aliensboot-server/module/passport"
	passportdb "github.com/KylinHe/aliensboot-server/module/passport/db"
	"time"
)

func init() {
	go func() {
		aliensboot.Run(
			database.Module,
			oplog.Module,
			passport.Module,
			defaultmodule.Module,
			gate.Module,
			gate.SubModule,
		)
	}()

	time.Sleep(3*time.Second)
	CleanDB()
}

// 清除数据库
func CleanDB() {
	//_ = socialdb.Database.DropCollections()
	passportdb.TestDropCollections()
}
