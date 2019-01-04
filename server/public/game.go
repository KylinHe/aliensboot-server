package main

import (
	"github.com/KylinHe/aliensboot-core"
	"github.com/KylinHe/aliensboot-core/module/database"
	"github.com/KylinHe/aliensboot-server/module/game"
)

func init() {

}

func main() {
	aliens.Run(
		database.Module,
		game.Module,
	)
}
