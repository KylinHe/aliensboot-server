/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/3/28
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package network

import (
	"github.com/KylinHe/aliensboot-core/cluster/center"
	"github.com/KylinHe/aliensboot-core/common/util"
	"github.com/KylinHe/aliensboot-core/gate"
	"github.com/KylinHe/aliensboot-core/log"
	"github.com/KylinHe/aliensboot-core/protocol/base"
	"github.com/KylinHe/aliensboot-server/constant"
	"github.com/KylinHe/aliensboot-server/dispatch"
	"github.com/KylinHe/aliensboot-server/module/gate/conf"
	"github.com/KylinHe/aliensboot-server/module/gate/route"
	"github.com/KylinHe/aliensboot-server/protocol"
	"net"
	"time"
)

func NewNetwork(agent gate.Agent) *Network {
	network := &Network{agent: agent, createTime: time.Now(), heartbeatTime: time.Now()}
	network.hashKey = agent.RemoteAddr().String()
	network.bindRoutes = make(map[uint16]string)
	network.bindServices = make(map[string]string)
	return network
}

type Network struct {
	agent gate.Agent
	//channel       chan *base.Any //消息管道

	authID int64 //用户标识 登录验证后

	hashKey string //用来做一致性负载均衡的标识

	createTime time.Time //创建时间

	heartbeatTime time.Time //上次的心跳时间

	bindRoutes map[uint16]string //绑定路由表 对应服务消息转发到指定节点上 比如场景服务器需要固定转发服务器

	//绑定的服务 离线需要通知
	bindServices map[string]string
}

//发送消息给客户端
func (this *Network) Push(msg *base.Any) {
	this.agent.WriteMsg(msg)
}

func (this *Network) KickOut(kickType protocol.KickType) {
	pushMsg := &protocol.Response{
		Gate: &protocol.Response_Kick{
			Kick: kickType,
		},
	}
	data, _ := pushMsg.Marshal()
	this.Push(&base.Any{Id: 1000, Value: data})
	this.agent.Close()
}

func (this *Network) OnClose() {
	if !this.IsAuth() {
		return
	}
	offlineMsg := &base.Any{
		Id:     constant.MsgOffline,
		AuthId: this.authID,
	}
	for service, node := range this.bindServices {
		dispatch.RequestNode(service, node, offlineMsg)
	}
}

//func (this *Network) requestCallback(request *base.Any, err error) {
//	Manager.acceptResponse(this, request, err)
//}

func (this *Network) handleResponse(response *base.Any, err error) {
	if response == nil || response.Value == nil || len(response.Value) == 0 {
		return
	}
	//TODO 返回服务不可用,或者尝试重发其他有效的节点
	if err != nil {
		log.Debugf("handle response %v err : %v", response, err)
		return
	}
	//更新验权id
	if response.GetAuthId() > 0 && !this.IsAuth() {
		this.Auth(response.GetAuthId())
	}
	if response.GetValue() == nil || len(response.GetValue()) == 0 {
		return
	}
	this.agent.WriteMsg(response)
	//lpc.StatisticsHandler.AddServiceStatistic(route.GetServiceByeSeq(response.Id), 1, 0.001)
}

func (this *Network) HandleMessage(request *base.Any) {
	//根据是否授权，传递授权id
	if this.IsAuth() {
		request.AuthId = this.authID
	} else {
		request.AuthId = 0
	}

	//消息增加网关id
	request.GateId = center.ClusterCenter.GetNodeID()
	node, _ := this.bindRoutes[request.Id]

	//var err error = nil
	//if node != "" {
	//	err = route.AsyncHandleNodeMessage(node, service.NewAsyncCall(request, handler.Go, this.handleResponse))
	//} else {
	//	err = route.AsyncHandleMessage(this.hashKey, service.NewAsyncCall(request, handler.Go, this.handleResponse))
	//}
	//if err != nil {
	//	log.Debug(err.Error())
	//}

	var response *base.Any = nil
	var err error = nil
	if node != "" {
		response, err = route.HandleNodeMessage(request, node)
	} else {
		response, err = route.HandleMessage(request, this.hashKey)
	}

	req := &protocol.Request{}
	req.Unmarshal(request.GetValue())

	resp := &protocol.Response{}
	resp.Unmarshal(response.GetValue())

	log.Debugf("request %+v : response:%+v", req, resp)

	this.handleResponse(response, err)

	////更新验权id
	//if response.GetAuthId() > 0 {
	//	this.Auth(response.GetAuthId())
	//}
	//if response != nil {
	//	this.agent.WriteMsg(response)
	//}
	//lpc.StatisticsHandler.AddServiceStatistic("passport", int32(request.Id), time.Now().Sub(start).Seconds())
}

//绑定服务节点,固定转发
func (this *Network) BindService(binds map[string]string) {
	for serviceName, serviceID := range binds {
		serviceSeq := route.GetServiceSeq(serviceName)
		if serviceSeq <= 0 {
			log.Errorf("bind service node error , service %v seq not found", serviceName)
			continue
		}
		this.bindRoutes[serviceSeq] = serviceID
	}
	this.bindServices = binds
}

func (this *Network) GetRemoteAddr() net.Addr {
	return this.agent.RemoteAddr()
}

func (this *Network) IsAuth() bool {
	return this.authID != 0
}

func (this *Network) Auth(id int64) {
	this.authID = id
	this.hashKey = util.Int64ToString(id)
	Manager.auth(id, this)
	//Skeleton.ChanRPCServer.Go(CommandAgentAuth, id, this)
}

//是否没有验权超时 释放多余的空连接
func (this *Network) IsAuthTimeout() bool {
	return !this.IsAuth() && time.Now().Sub(this.createTime).Seconds() >= conf.Config.AuthTimeout
}

//是否心跳超时
func (this *Network) IsHeartbeatTimeout() bool {
	return time.Now().Sub(this.heartbeatTime).Seconds() >= conf.Config.HeartbeatTimeout
}

func (this *Network) HeartBeat() {
	this.heartbeatTime = time.Now()
}
