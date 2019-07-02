package internal

import (
	"github.com/KylinHe/aliensboot-core/gate"
	"github.com/KylinHe/aliensboot-server/module/gate/conf"
	"github.com/KylinHe/aliensboot-server/module/gate/msg"
)

type SubModule struct {
	*gate.Gate
}

func (m *SubModule) GetName() string {
	return "gate"
}

func (m *SubModule) GetConfig() interface{} {
	return &conf.Config
}

func (m *SubModule) OnInit() {
	m.Gate = &gate.Gate{
		TcpConfig:    conf.Config.TCP,
		KcpConfig:    conf.Config.KCP,
		WsConfig:     conf.Config.WebSocket,
		Processor:    msg.Processor,
		AgentChanRPC: Skeleton.ChanRPCServer,
	}
}

func (m *SubModule) OnDestroy() {

}