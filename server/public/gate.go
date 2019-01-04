package main

import (
	"github.com/KylinHe/aliensboot-core"
	"github.com/KylinHe/aliensboot-server/module/gate"
)

func init() {

}

func main() {

	aliens.Run(
		gate.Module,
	)

}
