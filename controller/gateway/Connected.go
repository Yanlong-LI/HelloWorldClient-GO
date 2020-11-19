package gateway

import (
	"fmt"
	"github.com/yanlong-li/hi-go-gateway/common"
	"github.com/yanlong-li/hi-go-gateway/packet_model"
	"github.com/yanlong-li/hi-go-gateway/packet_model/server"
	logger "github.com/yanlong-li/hi-go-logger"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
	"net"
)

func init() {
	route.Register(common.GatewayAndServerGroup, packet_model.Connected{}, Connected)
}

func Connected(conn connect.Connector) {
	logger.Debug("连接到网关", 0, conn.GetId())

	fmt.Println(net.InterfaceAddrs())

	a := net.ParseIP("127.0.0.1")

	_ = conn.Send(server.Register{Name: "s1", Version: "v1", IP: a, Port: 3002, PeakLoad: 50000, OptimumLoad: 10000, Weight: 1})
}
