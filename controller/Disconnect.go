package controller

import (
	"github.com/yanlong-li/hi-go-server/common"
	"github.com/yanlong-li/hi-go-server/packet_model"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
)

func init() {
	route.Register(packet_model.Disconnect{}, Disconnect)
}

func Disconnect(conn connect.Connector) {
	logger.Debug("一个连接断开:", 0, conn.GetId())
	common.SignOut(conn.GetId())
}
