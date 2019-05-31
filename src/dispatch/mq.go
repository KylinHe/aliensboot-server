/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/5/21
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package dispatch

//var MQ = newMQDispatcher(conf.Config.MQ)
//
//type Dispatcher interface {
//	//RegisterConsumer(consumerID string, handle func(data *protocol.Any) error) //注册消息消费者
//	//UNRegisterConsumer(consumerID string)  //注销消息消费者
//
//	AsyncBroadcast(serviceType string, message proto.Message) error
//	AsyncPush(serviceType string, serviceID string, message proto.Message) error //异步推送
//
//	SyncRequest(serviceType string, message proto.Message) (interface{}, error)
//	SyncRequestNode(serviceType string, serviceID string, message proto.Message) (interface{}, error)
//	Request(serviceType string, message interface{}) (interface{}, error)
//	RequestNode(serviceType string, serviceID string, message interface{}) (interface{}, error)
//}
//
//func newMQDispatcher(config mq.Config) *MQDispatcher {
//	log.Debugf("%v", config.Address)
//	producer, err := mq.NewProducer(config)
//	if err != nil {
//		log.Fatal(err)
//	} else {
//		log.Info("init message producer success")
//	}
//	return &MQDispatcher{config: config, producer: producer, consumer: make(map[string]mq.IConsumer)}
//}
//
//type MQDispatcher struct {
//	config mq.Config
//	producer mq.IProducer
//	consumer map[string]mq.IConsumer
//}
//
////注册消息队列消费者 一般用来处理推送消息
//func (dispatcher *MQDispatcher) RegisterConsumer(serviceType string, handle func(data *protocol.Any) error) {
//	consumerID := serviceType + center.ClusterCenter.GetNodeID()
//	consumer := dispatcher.consumer[consumerID]
//	if consumer != nil {
//		log.Warnf("consumer %v already register", consumerID)
//		return
//	}
//
//	handleProxy := NewProtobufHandler(handle).Request
//	consumer, err := mq.NewConsumer(dispatcher.config, serviceType, center.ClusterCenter.GetNodeID(), handleProxy)
//	if err != nil {
//		log.Fatal(err)
//	} else {
//		dispatcher.consumer[consumerID] = consumer
//		log.Infof("init consumer %v success", consumerID)
//	}
//}
//
////注销消费者
//func (dispatcher *MQDispatcher) UNRegisterConsumer(serviceType string) {
//	consumerID := serviceType + center.ClusterCenter.GetNodeID()
//	consumer := dispatcher.consumer[consumerID]
//	if consumer != nil {
//		err := consumer.Close()
//		if err != nil {
//			log.Error(err)
//		}
//	}
//}
//
//func NewProtobufHandler(proxy func(message *protocol.Any) error) *protobufHandler {
//	return &protobufHandler{proxy}
//}
//
//type protobufHandler struct {
//	proxy func(message *protocol.Any) error
//}
//
//func (this *protobufHandler) Request(data []byte) error {
//	requestProxy := &protocol.Any{}
//	error := proto.Unmarshal(data, requestProxy)
//	if error != nil {
//		return error
//	}
//	return this.proxy(requestProxy)
//}
//
////网关异步推送信息给指定用户
//func (dispatcher *MQDispatcher) GatePush(serviceType string, authID int64, message proto.Message) error {
//	data, err := proto.Marshal(message)
//	if err != nil {
//		return err
//	}
//	request := &protocol.Any{TypeUrl: serviceType, Value: data}
//	//gateID := getGateNodeID(clientID)
//	gateID := cache.ClusterCache.GetAuthGateID(authID)
//	if gateID == "" {
//		return errors.New(fmt.Sprintf("gate ID can not found, clientID : %v", gateID))
//	}
//	return dispatcher.AsyncPush(constant.SERVICE_GATE, gateID, request)
//}
//
////func getGateNodeID(clientID string) string {
////	return strings.Split(clientID, "_")[0]
////}
//
////消息异步推送 - 推送到指定服务节点
//func (dispatcher *MQDispatcher) AsyncPush(serviceType string, serviceID string, message proto.Message) error {
//	data, err := proto.Marshal(message)
//	if err != nil {
//		return err
//	}
//	dispatcher.producer.Push(serviceType, serviceID, data)
//	return nil
//}
//
////消息异步推送 - 广播所有对应服务节点
//func (dispatcher *MQDispatcher) AsyncBroadcast(serviceType string, message proto.Message) error {
//	data, err := proto.Marshal(message)
//	if err != nil {
//		return err
//	}
//	dispatcher.producer.Broadcast(serviceType, data)
//	return nil
//}
