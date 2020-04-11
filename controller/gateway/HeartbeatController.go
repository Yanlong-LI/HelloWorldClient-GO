package gateway

import (
	"fmt"
	"github.com/yanlong-li/HelloWorld-GO/io/logger"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorldServer/packetModel/gateway"
)

func init() {
	route.Register(gateway.Heartbeat{}, Heartbeat)
}

func Heartbeat(heartbeat gateway.Heartbeat, conn connect.Connector) {
	logger.Debug(fmt.Sprintf("收到 {%d} 心跳包 {%d}", conn.GetId(), heartbeat.Sn), 0, heartbeat)
	_ = conn.Send(gateway.Heartbeat{Sn: heartbeat.Sn + 1})
}
