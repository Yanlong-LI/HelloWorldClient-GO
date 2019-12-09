package gateway

import (
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/route"
	"HelloWorldServer/packet/gateway"
)

func init() {
	route.Register(gateway.Heartbeat{}, Heartbeat)
}

func Heartbeat(heartbeat gateway.Heartbeat, conn connect.Connector) {

	conn.Send(gateway.Heartbeat{Sn: heartbeat.Sn + 1})
}
