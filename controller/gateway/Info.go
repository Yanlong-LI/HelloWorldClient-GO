package gateway

import (
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/route"
	"HelloWorldServer/packet/gateway"
)

func init() {
	route.Register(gateway.Info{}, GetInfo)
}

func GetInfo(info gateway.Info, conn connect.Connector) {

	conn.Send(gateway.Info{Name: "Master Service", Version: "1.0.0", Region: "China", CreateTime: 1575703496})
}
