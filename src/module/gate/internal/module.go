package internal

import (
	"github.com/KylinHe/aliensboot-server/module/gate/cache"
	"github.com/KylinHe/aliensboot-server/module/gate/conf"
	"github.com/KylinHe/aliensboot-server/module/gate/network"
	"github.com/KylinHe/aliensboot-server/module/gate/route"
	"github.com/KylinHe/aliensboot-server/module/gate/service"
	"github.com/KylinHe/aliensboot-core/module/base"
)

var Skeleton = base.NewSkeleton()

type Module struct {
	*base.Skeleton
}

func (m *Module) GetName() string {
	return "gate"
}

func (m *Module) GetConfig() interface{} {
	return &conf.Config
}

func (m *Module) OnInit() {
	m.Skeleton = Skeleton
	route.Init()
	cache.Init()
	network.Init(Skeleton)
	service.Init(Skeleton.ChanRPCServer)
	//http.Init(conf.Config.Http)
}

func (m *Module) OnDestroy() {
	//http.Close()
	service.Close()
}
