package controller

import (
	"HelloWorld/io/network/packet"
	"HelloWorld/io/network/route"
	"HelloWorld/io/network/socket/connect"
	"HelloWorldServer/data/user"
	"fmt"
	"time"
)

func init() {

	route.Register(packet.Login{}, Login)
}
func Login(test packet.Login, conn *connect.Connector) {
	fmt.Println("用户登录：", test.Username)

	if test.Username == "张三" {
		conn.Send(packet.Token{Token: "123456789123456789"})
		user.Register(conn.ID, user.User{Name: test.Username, Time: time.Now()})
		fmt.Println(test.Username, "登录成功")
	} else {
		conn.Send(packet.LoginFail{Code: -1, Message: "用户名不正确"})
		fmt.Println("用户名不正确，已拒绝登录")
	}
}
