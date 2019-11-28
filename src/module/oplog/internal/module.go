package internal

import (
	"github.com/KylinHe/aliensboot-server/module/oplog/cache"
	"github.com/KylinHe/aliensboot-server/module/oplog/conf"
	"github.com/KylinHe/aliensboot-server/module/oplog/db"
	"github.com/KylinHe/aliensboot-core/module/base"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*base.Skeleton
}

func (m *Module) GetName() string {
	return "oplog"
}

func (m *Module) GetConfig() interface{} {
	return &conf.Config
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	cache.Init()
	//conf.Init()
	db.Init()
}

func (m *Module) OnDestroy() {
	db.Close()
	cache.Close()
	//conf.Close()
}
