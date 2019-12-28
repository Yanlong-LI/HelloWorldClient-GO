package controller

import (
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/route"
	"HelloWorldServer/packet"
	"fmt"
)

func init() {
	route.Register(packet.Connected{}, Connected)
}

func Connected(conn connect.Connector) {

	fmt.Println("新连接", conn.GetId())

	//conn.Send(gateway.Info{Name: "Master Service", Version: "1.0.0", Region: "China", CreateTime: 1575703496})
	//conn.Send(gateway.GetHeartbeat{})
}
