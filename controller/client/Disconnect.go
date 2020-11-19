package client

import (
	"github.com/yanlong-li/hi-go-gateway/common"
	"github.com/yanlong-li/hi-go-gateway/packet_model"
	logger "github.com/yanlong-li/hi-go-logger"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
)

func init() {
	route.Register(common.ServerAndClientGroup, packet_model.Disconnect{}, Disconnect)
}

func Disconnect(conn connect.Connector) {
	logger.Debug("一个连接断开:", 0, conn.GetId())
}
