package main

import (
	"github.com/KylinHe/aliensboot-core/aliensboot"
	"github.com/KylinHe/aliensboot-core/module/database"
	"github.com/KylinHe/aliensboot-server/module/gate"
	"github.com/KylinHe/aliensboot-server/module/oplog"
	"github.com/KylinHe/aliensboot-server/module/passport"
)

func init() {

}

func main() {

	aliensboot.Run(
		database.Module,
		oplog.Module,
		passport.Module,
		gate.Module,
		gate.SubModule,
	)

}
