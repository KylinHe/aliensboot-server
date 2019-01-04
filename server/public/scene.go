package main

import (
	"github.com/KylinHe/aliensboot-core"
	"github.com/KylinHe/aliensboot-core/module/database"
	"github.com/KylinHe/aliensboot-server/module/game"
	"github.com/KylinHe/aliensboot-server/module/gate"
	"github.com/KylinHe/aliensboot-server/module/passport"
	"github.com/KylinHe/aliensboot-server/module/scene"
)

func init() {

}

func main() {

	aliens.Run(
		database.Module,
		gate.Module,
		passport.Module,
		game.Module,
		scene.Module,
	)

}
