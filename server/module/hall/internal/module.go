package internal

import (
	"github.com/KylinHe/aliensboot-core/module/base"
	"github.com/KylinHe/aliensboot-server/module/hall/cache"
	"github.com/KylinHe/aliensboot-server/module/hall/conf"
	"github.com/KylinHe/aliensboot-server/module/hall/db"
	"github.com/KylinHe/aliensboot-server/module/hall/service"
	"github.com/KylinHe/aliensboot-server/module/hall/task"
)

type Module struct {
	*base.Skeleton
}

func (m *Module) GetName() string {
	return "hall"
}

func (m *Module) GetConfig() interface{} {
	return &conf.Config
}

func (m *Module) OnInit() {
	m.Skeleton = base.NewSkeleton()
	conf.Init()
	db.Init()
	cache.Init()
	task.Init(m.Skeleton)
	service.Init(m.ChanRPCServer)
}

func (m *Module) OnDestroy() {
	service.Close()
	cache.Close()
	db.Close()
	conf.Close()
}
