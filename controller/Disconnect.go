package controller

import (
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/route"
	"HelloWorldServer/data/user"
	"HelloWorldServer/packet"
	"fmt"
)

func init() {
	route.Register(packet.Disconnect{}, Disconnect)
}

func Disconnect(ID uint32) {
	fmt.Println("一个连接断开:", ID)
	userInfo, ok := user.GetUser(ID)
	if ok {
		userPacket := packet.User{Nickname: userInfo.Name, LoginTime: userInfo.Time.Unix()}
		fmt.Println("用户注销：", userInfo.Name)
		user.RemoveUser(ID)
		connect.Broadcast(packet.Logout{User: userPacket})
	}
}
