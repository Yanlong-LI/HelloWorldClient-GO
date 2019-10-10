package controller

import (
	"HelloWorld/io/network/route"
	"HelloWorld/io/network/socket/connect"
	"HelloWorldServer/data/user"
	"HelloWorldServer/packet"
	"fmt"
)

func init() {
	route.Register(packet.Disconnect{}, Disconnect)
}

func Disconnect(disconnect packet.Disconnect, ID uint32) {

	u := user.GetUser(ID)
	user2 := packet.User{Nickname: u.Name, LoginTime: u.Time.Unix()}
	fmt.Println("用户注销：", u.Name)
	user.RemoveUser(ID)
	connect.Broadcast(packet.Logout{User: user2})
}
