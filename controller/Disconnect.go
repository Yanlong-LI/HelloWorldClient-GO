package controller

import (
	"github.com/yanlong-li/HelloWorld-GO/io/logger"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorldServer/model/online"
	"github.com/yanlong-li/HelloWorldServer/packetModel"
)

func init() {
	route.Register(packetModel.Disconnect{}, Disconnect)
}

func Disconnect(conn connect.Connector) {
	logger.Debug("一个连接断开:", 0, conn.GetId())
	online.SignOut(conn.GetId())
}
