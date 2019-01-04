package internal

import (
	"github.com/KylinHe/aliensboot-core/module/base"
	"github.com/KylinHe/aliensboot-server/module/room/cache"
	"github.com/KylinHe/aliensboot-server/module/room/conf"
	"github.com/KylinHe/aliensboot-server/module/room/db"
	"github.com/KylinHe/aliensboot-server/module/room/service"
)

type Module struct {
	*base.Skeleton
}

func (m *Module) GetName() string {
	return "room"
}

func (m *Module) GetConfig() interface{} {
	return &conf.Config
}

func (m *Module) OnInit() {
	m.Skeleton = base.NewSkeleton()
	conf.Init()
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
