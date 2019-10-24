package controller

import (
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/route"
	"HelloWorldServer/data/user"
	"HelloWorldServer/packet"
	"fmt"
)

func init() {
	route.Register(packet.Message{}, Message)
}

func Message(message packet.Message, conn connect.Connector) {
	userInfo, ok := user.GetUser(conn.GetId())
	if ok {
		fmt.Println(userInfo.Name, ":", message.Content)
		userPacket := packet.User{Nickname: userInfo.Name, LoginTime: userInfo.Time.Unix()}
		connect.Broadcast(packet.GlobalMessage{Message: message, User: userPacket})
	}
}
