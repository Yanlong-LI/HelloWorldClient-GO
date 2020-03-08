package gateway

import (
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorldServer/packetModel/gateway"
	"github.com/yanlong-li/HelloWorldServer/packetModel/user"
)

func init() {
	route.Register(gateway.Heartbeat{}, Heartbeat)
}

func Heartbeat(heartbeat gateway.Heartbeat, conn connect.Connector) {

	conn.Send(gateway.Heartbeat{Sn: heartbeat.Sn + 1})
	conn.Broadcast(user.Info{Id: conn.GetId(), Nickname: "test"}, true)
}
