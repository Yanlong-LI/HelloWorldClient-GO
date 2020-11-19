package client

import (
	"github.com/yanlong-li/hi-go-gateway/common"
	logger "github.com/yanlong-li/hi-go-logger"
	"github.com/yanlong-li/hi-go-server/packet_model"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
)

func init() {
	route.Register(common.ServerAndClientGroup, packet_model.Disconnect{}, Disconnect)
}

func Disconnect(conn connect.Connector) {
	logger.Debug("一个连接断开:", 0, conn.GetId())
}
func init() {
	route.Register(common.ServerAndClientGroup, packet_model.Connected{}, Connected)
}

func Connected(conn connect.Connector) {
	logger.Debug("新连接", 0, conn.GetId())
}
