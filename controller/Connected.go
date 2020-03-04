package controller

import (
	"fmt"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorldServer/packet"
)

func init() {
	route.Register(packet.Connected{}, Connected)
}

func Connected(conn connect.Connector) {

	fmt.Println("新连接", conn.GetId())

	//conn.Send(gateway.Info{Name: "Master Service", Version: "1.0.0", Region: "China", CreateTime: 1575703496})
	//conn.Send(gateway.GetHeartbeat{})
}
