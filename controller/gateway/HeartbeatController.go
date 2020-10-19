package gateway

import (
	"fmt"
	logger "github.com/yanlong-li/hi-go-logger"
	"github.com/yanlong-li/hi-go-server/packet_model/gateway"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
)

func init() {
	route.Register(gateway.Heartbeat{}, Heartbeat)
}

func Heartbeat(heartbeat gateway.Heartbeat, conn connect.Connector) {
	logger.Debug(fmt.Sprintf("收到 {%d} 心跳包 {%d}", conn.GetId(), heartbeat.Sn), 0, heartbeat)
	_ = conn.Send(gateway.Heartbeat{Sn: heartbeat.Sn + 1})
}
