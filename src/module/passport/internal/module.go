package internal

import (
	"github.com/KylinHe/aliensboot-server/module/passport/conf"
	"github.com/KylinHe/aliensboot-server/module/passport/db"
	"github.com/KylinHe/aliensboot-server/module/passport/service"
	"github.com/KylinHe/aliensboot-core/module/base"
)

type Module struct {
	*base.Skeleton
}

func (m *Module) GetName() string {
	return "passport"
}

func (m *Module) GetConfig() interface{} {
	return &conf.Config
}

func (m *Module) OnInit() {
	m.Skeleton = base.NewSkeleton()
	conf.InitConfig()
	conf.Init()
	db.Init()
	service.Init(m.ChanRPCServer)
}

func (m *Module) OnDestroy() {
	service.Close()
	db.Close()
}
