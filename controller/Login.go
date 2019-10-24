package controller

import (
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/route"
	"HelloWorldServer/data/user"
	"HelloWorldServer/packet"
	"fmt"
	"time"
)

func init() {

	route.Register(packet.Login{}, Login)
}
func Login(test packet.Login, conn connect.Connector) {
	fmt.Println("用户登录：", test.Username)

	if test.Username != "" {
		conn.Send(packet.Token{Token: "123456789123456789"})
		user.Register(user.User{Id: conn.GetId(), Name: test.Username, Time: time.Now()})
		fmt.Println(test.Username, "登录成功")
		userPacket := packet.User{Nickname: test.Username, LoginTime: time.Now().Unix()}
		connect.Broadcast(packet.NewUser{User: userPacket})
	} else {
		conn.Send(packet.LoginFail{Code: -1, Message: "用户名不正确"})
		fmt.Println("用户名不正确，已拒绝登录")
	}
}
