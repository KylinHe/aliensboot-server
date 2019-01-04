package main

import (
	"github.com/KylinHe/aliensboot-core/aliensboot"
	"github.com/KylinHe/aliensboot-core/module/database"
	"github.com/KylinHe/aliensboot-server/module/defaultmodule"
)

func init() {

}

func main() {

	aliensboot.Run(
		database.Module,
		defaultmodule.Module,
	)
}
