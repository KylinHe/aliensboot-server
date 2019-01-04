package internal

import (
	"github.com/KylinHe/aliensboot-core/gate"
	"github.com/KylinHe/aliensboot-core/module/base"
	"github.com/KylinHe/aliensboot-server/module/gate/cache"
	"github.com/KylinHe/aliensboot-server/module/gate/conf"
	"github.com/KylinHe/aliensboot-server/module/gate/http"
	"github.com/KylinHe/aliensboot-server/module/gate/msg"
	"github.com/KylinHe/aliensboot-server/module/gate/network"
	"github.com/KylinHe/aliensboot-server/module/gate/route"
	"github.com/KylinHe/aliensboot-server/module/gate/service"
)

var Skeleton = base.NewSkeleton()

type Module struct {
	*gate.Gate
}

func (m *Module) GetName() string {
	return "gate"
}

func (m *Module) GetConfig() interface{} {
	return &conf.Config
}

func (m *Module) OnInit() {
	//conf.Init(m.GetName())
	m.Gate = &gate.Gate{
		TcpConfig:    conf.Config.TCP,
		WsConfig:     conf.Config.WebSocket,
		Processor:    msg.Processor,
		AgentChanRPC: Skeleton.ChanRPCServer,
	}
	route.Init()
	cache.Init()
	network.Init(Skeleton)
	service.Init(Skeleton.ChanRPCServer)
	http.Init(conf.Config.Http)
}

func (m *Module) OnDestroy() {
	http.Close()
	service.Close()
}

func (m *Module) Run(closeSig chan bool) {
	go m.Gate.Run(closeSig)
	Skeleton.Run(closeSig)
}
