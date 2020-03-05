package controller

import (
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	conn2 "github.com/yanlong-li/HelloWorldServer/model/online"
	"github.com/yanlong-li/HelloWorldServer/packetModel"
	"github.com/yanlong-li/HelloWorldServer/packetModel/gateway"
	"github.com/yanlong-li/HelloWorldServer/packetModel/trait"
)

var WhiteList = make(map[uint32]bool, 1)

func init() {
	route.Register(packetModel.BeforeRecving{}, Middleware)

	WhiteList[7001] = true
	WhiteList[6001] = true
	WhiteList[6007] = true
	WhiteList[6010] = true
	WhiteList[6013] = true
}

func Middleware(OpCode uint32, data []byte, conn connect.Connector) bool {
	if _, ok := WhiteList[OpCode]; !ok {
		// 验证用户是否登陆
		_, err := conn2.Auth(conn.GetId())
		if err != nil {
			conn.Send(gateway.AuthenticFail{Fail: trait.Fail{Code: 7015, Message: "当前未登陆"}})
			return false
		}

	}

	return true
}
