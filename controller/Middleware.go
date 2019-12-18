package controller

import (
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/route"
	conn2 "HelloWorldServer/model/Login"
	"HelloWorldServer/packet"
	"HelloWorldServer/packet/gateway"
	"HelloWorldServer/packet/trait"
)

var WhiteList = make(map[uint32]bool, 1)

func init() {
	route.Register(packet.RecvPacketMiddleware{}, Middleware)

	WhiteList[7001] = true
	WhiteList[6001] = true
	WhiteList[6007] = true
}

func Middleware(OpCode uint32, conn connect.Connector) bool {
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
