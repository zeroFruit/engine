package p2p

//todo implement it!
func init() {
	//publisher := rabbitmq.Connect("")
	//eventRepository := midgard.NewRepo(store, publisher)
	//publisher :=
	//grpcCommandService := adapter.NewGrpcCommandService(publisher)
	//peerApi := api.NewPeerApi(eventRepository, grpcCommandService)
	////myIp := conf.GetConfiguration().Common.NodeIp
	//
	//config := conf.GetConfiguration()
	////create rabbitmq client
	//rabbitmqClient := rabbitmq.Connect(config.Common.Messaging.Url)
	//// todo change node repo and leader repo after managing peerTable
	//
	//messageDispatcher := adapter.NewMessageDispatcher(publisher)
	//
	//
	////create amqp Handler
	//eventHandler := adapter.NewNodeEventHandler(nodeRepository, leaderRepository)
	//grpcMessageHandler := adapter.NewGrpcMessageHandler(leaderApi, nodeApi, messageDispatcher)
	//
	//// Subscribe amqp server
	//err1 := rabbitmqClient.Subscribe("Command", "connection.*", eventHandler)
	//
	//if err1 != nil {
	//	panic(err1)
	//}
	//
	//err2 := rabbitmqClient.Subscribe("Command", "message.*", grpcMessageHandler)
	//
	//if err2 != nil {
	//	panic(err2)
	//}
}
