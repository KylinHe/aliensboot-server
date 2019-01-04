package network

import (
	"github.com/KylinHe/aliensboot-core/cluster/center"
	"github.com/KylinHe/aliensboot-core/common/data_structures/set"
	"github.com/KylinHe/aliensboot-core/common/util"
	"github.com/KylinHe/aliensboot-core/log"
	modulebase "github.com/KylinHe/aliensboot-core/module/base"
	"github.com/KylinHe/aliensboot-core/protocol/base"
	"github.com/KylinHe/aliensboot-server/dispatch/rpc"
	"github.com/KylinHe/aliensboot-server/module/gate/cache"
	"github.com/KylinHe/aliensboot-server/protocol"
)

var Manager = &networkManager{}

//const (
//	CommandRpcResponse = "resp"
//)

type networkManager struct {
	//handler *modulebase.Skeleton
	networks     *set.HashSet       //存储所有未验权的网络连接
	authNetworks map[int64]*Network //存储所有验权通过的网络连接
	node         string             //当前节点名
	//timeWheel *util.TimeWheel       //验权检查时间轮
}

var handler *modulebase.Skeleton

func Init(skeleton *modulebase.Skeleton) {
	handler = skeleton
	Manager.Init()
}

//开启权限,心跳等验证机制
func (this *networkManager) Init() {
	///this.handler = chanRpc
	//this.chanRpc.Register(CommandRpcResponse, this.handleResponse)
	this.node = center.ClusterCenter.GetNodeID()
	this.networks = set.NewHashSet()
	this.authNetworks = make(map[int64]*Network)

	//心跳精确到s
	//this.timeWheel = util.NewTimeWheel(time.Second, 60, this.dealAuthTimeout)
	//this.timeWheel.Start()
}

//
//func (this *networkManager) AsyncCall(f func(), c func()) {
//	this.handler.Go(f, c)
//}

//func (this *networkManager) handleResponse(args []interface{}) {
//	network := args[0].(*Network)
//	response := args[1].(*base.Any)
//	err, ok := args[2].(error)
//	if ok {
//		network.handleResponse(response, err)
//	} else {
//		network.handleResponse(response, nil)
//	}
//}

func (this *networkManager) BindService(authID int64, binds map[string]string) {
	auth := this.authNetworks[authID]
	if auth == nil {
		return
	}
	auth.BindService(binds)
}

func (this *networkManager) KickOut(authID int64, kickType protocol.KickType) {
	auth := this.authNetworks[authID]
	if auth == nil {
		return
	}
	auth.KickOut(kickType)
}

//推送消息
func (this *networkManager) Push(authID int64, message *base.Any) {
	auth := this.authNetworks[authID]
	if auth == nil {
		log.Debugf("auth network not found %v", authID)
		return
	}
	auth.Push(message)
}

//广播消息
func (this *networkManager) Broadcast(message *base.Any) {
	for _, network := range this.authNetworks {
		network.Push(message)
	}
}

func (this *networkManager) AddNetwork(network *Network) {
	data := make(util.TaskData)
	data[0] = network
	//this.timeWheel.AddTimer(time.Duration(conf.Config.AuthTimeout)*time.Second, network, data)
	this.networks.Add(network)
}

func (this *networkManager) RemoveNetwork(network *Network) {
	network.OnClose()
	if network.IsAuth() {
		storeNetwork := this.authNetworks[network.authID]
		if storeNetwork != nil && storeNetwork == network {
			delete(this.authNetworks, network.authID)
			cache.GateCache.CleanAuthGateID(network.authID)
		}
	} else {
		//this.timeWheel.RemoveTimer(network)
		this.networks.Remove(network)
	}

}

func (this *networkManager) DealAuthTimeout() {
	//this.networks.Range(func(element interface{}) {
	//	network := element.(*Network)
	//	//连接超过固定时长没有验证权限需要退出
	//	if network.IsAuthTimeout() {
	//		//log.Debug("Network auth timeout : %v", networker.GetRemoteAddr())
	//		network.KickOut(protocol.KickType_Timeout)
	//		this.networks.Remove(network)
	//	}
	//})
}

//验权限
func (this *networkManager) auth(authID int64, network *Network) {
	//this.timeWheel.RemoveTimer(network)
	this.networks.Remove(network)

	oldNetwork, ok := this.authNetworks[authID]

	this.authNetworks[authID] = network
	cache.GateCache.SetAuthGateID(authID, this.node)
	//顶号处理
	if ok {
		oldNetwork.KickOut(protocol.KickType_OtherSession)
	} else {
		node := cache.GateCache.GetAuthGateID(authID)
		//用户在其他网关节点登录 需要发送远程T人
		if node != "" && node != this.node {
			kickMsg := &protocol.KickOut{
				AuthID:   authID,
				KickType: protocol.KickType_OtherSession,
			}
			rpc.Gate.KickOut(node, kickMsg)
		}
	}


	//cache.ClusterCache.SetAuthGateID(authID, center.ClusterCenter.GetNodeID())
}
