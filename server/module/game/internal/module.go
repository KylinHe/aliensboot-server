package internal

import (
	"github.com/KylinHe/aliensboot-core/module/base"
	"github.com/KylinHe/aliensboot-server/module/game/conf"
	"github.com/KylinHe/aliensboot-server/module/game/db"
	"github.com/KylinHe/aliensboot-server/module/game/service"
)

type Module struct {
	*base.Skeleton
}

func (m *Module) GetName() string {
	return "game"
}

func (m *Module) GetConfig() interface{} {
	return &conf.Config
}

func (m *Module) OnInit() {
	m.Skeleton = base.NewSkeleton()
	db.Init()
	service.Init(m.ChanRPCServer)
}

func (m *Module) OnDestroy() {
	service.Close()
	db.Close()
}
