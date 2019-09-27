package controller

import (
	"HelloWorld/io/network/packet"
	"HelloWorld/io/network/route"
	"HelloWorld/io/network/socket/connect"
	"fmt"
)

func init() {

	route.Register(packet.Login{}, Login)
}
func Login(test packet.Login, conn *connect.Connector) {
	fmt.Println("用户登录：", test.Username)

	if test.Username == "张三" {
		conn.Send(packet.Token{Token: "123456789123456789"})
	} else {
		conn.Send(packet.LoginFail{Code: 1, Message: "用户名不正确"})
	}

}
