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
	"github.com/KylinHe/aliensboot-server/constant"
	"github.com/KylinHe/aliensboot-server/dispatch"
	"github.com/KylinHe/aliensboot-server/module/gate/conf"
	"github.com/KylinHe/aliensboot-server/module/gate/route"
	"github.com/KylinHe/aliensboot-server/protocol"
	"github.com/KylinHe/aliensboot-core/cluster/center"
	"github.com/KylinHe/aliensboot-core/common/util"
	"github.com/KylinHe/aliensboot-core/gate"
	"github.com/KylinHe/aliensboot-core/log"
	"github.com/KylinHe/aliensboot-core/protocol/base"
	"github.com/KylinHe/aliensboot-core/task"
	"net"
	"time"
)

//心跳包
var heartbeat = &base.Any{Id: 0}

func NewNetwork(agent gate.Agent) *Network {
	network := &Network{agent: agent, createTime: time.Now(), heartbeatTime: time.Now()}
	network.hashKey = agent.RemoteAddr().String()
	network.bindRoutes = make(map[uint16]string)
	network.bindServices = make(map[string]string)

	//log.Debugf("%v", agent.RemoteAddr())
	//network.channel = make(chan *base.Any, 1000)
	return network
}

type Network struct {
	agent gate.Agent

	//channel chan *base.Any //消息管道

	authID int64 //用户标识 登录验证后

	hashKey string //用来做一致性负载均衡的标识

	createTime time.Time //创建时间

	heartbeatTime time.Time //上次的心跳时间

	bindRoutes map[uint16]string //绑定路由表 对应服务消息转发到指定节点上 比如需要固定转发到指定的场景服务器

	//绑定的服务 离线需要通知
	bindServices map[string]string //
}

func (this *Network) GetAddr() string {
	return this.agent.RemoteAddr().(*net.TCPAddr).IP.String()
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

	// 通知带会话的服务玩家下线
	for _, service := range conf.Config.Session {
		_ = dispatch.Send(service, offlineMsg, this.hashKey)
	}

	//for service, node := range this.bindServices {
	//	_, _ = dispatch.RequestNode(service, node, offlineMsg)
	//}
}

func (this *Network) HandleMessage(request *base.Any) {
	this.HeartBeat()
	//心跳包直接回
	if request.Id == 0 {
		this.agent.WriteMsg(heartbeat)
		return
	}

	//根据是否授权，传递授权id
	if this.IsAuth() {
		request.AuthId = this.authID
	} else {
		request.AuthId = 0
	}

	//消息增加网关id
	request.GateId = center.ClusterCenter.GetNodeID()
	node, _ := this.bindRoutes[request.Id]

	// 账号模块需要带上ip头信息
	if request.Id == constant.ModulePassportId {
		request.AddHeader(constant.HeaderIP, []byte(this.GetAddr()))
	}

	task.SafeGo(func() {
		var response *base.Any = nil
		var err error = nil
		if node != "" {
			response, err = route.HandleNodeMessage(request, node)
		} else {
			response, err = route.HandleMessage(request, this.hashKey)
		}
		//TODO 返回服务不可用,或者尝试重发其他有效的节点
		if err != nil {
			log.Debugf("handle response %v err : %v", response, err)
			return
		}
		if response == nil || response.Value == nil || len(response.Value) == 0 {
			return
		}
		response.Id = request.Id

		//req := &protocol.Request{}
		//req.Unmarshal(request.GetValue())
		//
		//resp := &protocol.Response{}
		//resp.Unmarshal(response.GetValue())

		//log.Debugf("auth %v request %+v : response:%+v",request.AuthId, req, resp)
		Manager.HandleResponse(this, response, err)
		//this.handleResponse(response, err)
	})
}

func (this *Network) handleResponse(response *base.Any, err error) {
	//更新验权id
	if response.GetAuthId() > 0 && !this.IsAuth() {
		this.Auth(response.GetAuthId())
	}
	this.agent.WriteMsg(response)
	//lpc.StatisticsHandler.AddServiceStatistic(route.GetServiceByeSeq(response.Id), 1, 0.001)
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
		this.bindServices[serviceName] = serviceID
	}
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

func (this *Network) GetAuthId() int64 {
	return this.authID
}

//是否没有验权超时 释放多余的空连接
func (this *Network) HandleAuthTimeout([]interface{}) {
	//!this.IsAuth() && time.Now().Sub(this.createTime).Seconds() >= conf.Config.AuthTimeout
	//未授权需要T除
	if !this.IsAuth() {
		this.KickOut(protocol.KickType_Timeout)
	}
}

//是否心跳超时
func (this *Network) HandleHeartbeatTimeout([]interface{}) {
	isTimeOut := time.Now().Sub(this.heartbeatTime).Seconds() >= conf.Config.HeartbeatTimeout
	if isTimeOut {
		this.KickOut(protocol.KickType_Timeout)
	}
}

func (this *Network) HeartBeat() {
	this.heartbeatTime = time.Now()
}
