package main

import (
	"github.com/KylinHe/aliensboot-core"
	"github.com/KylinHe/aliensboot-core/module/database"
	"github.com/KylinHe/aliensboot-server/module/hall"
)

func init() {

}

func main() {

	aliens.Run(
		database.Module,
		hall.Module,
	)
}
