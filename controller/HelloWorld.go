package controller

import (
	"HelloWorld/io/network/packet"
	"HelloWorld/io/network/route"
	"HelloWorld/io/network/websocket/connect"
	"fmt"
)

//func init() {
//	route.Register(packet.HelloWorld{}, HelloWorld)
//}
//
//func HelloWorld(world packet.HelloWorld, conn *connect.Connector) {
//	// 首次连接 服务端不做处理
//}

func init() {
	route.Register(packet.HelloWorld{}, HelloWorld)
}

func HelloWorld(world packet.HelloWorld, connector *connect.Connector) {
	fmt.Println(world)
}
