package internal

import (
	"github.com/KylinHe/aliensboot-core/module/base"
	"github.com/KylinHe/aliensboot-server/module/scene/cache"
	"github.com/KylinHe/aliensboot-server/module/scene/conf"
	"github.com/KylinHe/aliensboot-server/module/scene/db"
	"github.com/KylinHe/aliensboot-server/module/scene/handler"
	"github.com/KylinHe/aliensboot-server/module/scene/service"
)

type Module struct {
	*base.Skeleton
}

func (m *Module) GetName() string {
	return "scene"
}

func (m *Module) GetConfig() interface{} {
	return &conf.Config
}

func (m *Module) OnInit() {
	m.Skeleton = base.NewSkeleton()
	conf.Init()
	db.Init()
	cache.Init()
	handler.Init(m.Skeleton)
	service.Init(m.ChanRPCServer)
}

func (m *Module) OnDestroy() {
	service.Close()
	cache.Close()
	db.Close()
	conf.Close()
}
