package internal

import (
	"github.com/KylinHe/aliensboot-core/module/base"
	"github.com/KylinHe/aliensboot-server/module/defaultmodule/cache"
	"github.com/KylinHe/aliensboot-server/module/defaultmodule/conf"
	"github.com/KylinHe/aliensboot-server/module/defaultmodule/db"
	"github.com/KylinHe/aliensboot-server/module/defaultmodule/service"
)

type Module struct {
	*base.Skeleton
}

func (m *Module) GetName() string {
	return "defaultmodule"
}

func (m *Module) GetConfig() interface{} {
	return &conf.Config
}

func (m *Module) OnInit() {
	m.Skeleton = base.NewSkeleton()
	conf.Init(m.Skeleton)
	db.Init()
	cache.Init()
	service.Init(m.ChanRPCServer)
}

func (m *Module) OnDestroy() {
	service.Close()
	cache.Close()
	db.Close()
	conf.Close()
}
