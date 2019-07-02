package main

import (
	"github.com/KylinHe/aliensboot-core/aliensboot"
	"github.com/KylinHe/aliensboot-core/module/database"
	"github.com/KylinHe/aliensboot-server/module/gate"
	"github.com/KylinHe/aliensboot-server/module/hall"
	"github.com/KylinHe/aliensboot-server/module/passport"
	"github.com/KylinHe/aliensboot-server/module/room"
)

func init() {

}

func main() {

	aliensboot.Run(
		database.Module,
		hall.Module,
		room.Module,
		passport.Module,
		gate.Module,
		gate.SubModule,
		//game.Module,
		//scene.Module,
	)
}
