package service

import (
	"github.com/KylinHe/aliensboot-core/aliensboot"
	"github.com/KylinHe/aliensboot-core/module/database"
	"github.com/KylinHe/aliensboot-server/module/defaultmodule"
	"github.com/KylinHe/aliensboot-server/module/gate"
"github.com/KylinHe/aliensboot-server/module/oplog"
"github.com/KylinHe/aliensboot-server/module/passport"
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
}


