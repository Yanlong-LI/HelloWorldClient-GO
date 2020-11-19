package client

import (
	"github.com/yanlong-li/hi-go-gateway/common"
	"github.com/yanlong-li/hi-go-gateway/packet_model"
	logger "github.com/yanlong-li/hi-go-logger"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
)

func init() {
	route.Register(common.ServerAndClientGroup, packet_model.Connected{}, Connected)
}

func Connected(conn connect.Connector) {
	logger.Debug("新连接", 0, conn.GetId())
}
