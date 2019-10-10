package controller

import (
	"HelloWorld/io/network/route"
	"HelloWorld/io/network/socket/connect"
	user2 "HelloWorldServer/data/user"
	"HelloWorldServer/packet"
	"fmt"
)

func init() {
	route.Register(packet.Message{}, Message)
}

func Message(message packet.Message, conn *connect.Connector) {
	user := user2.GetUser(conn.ID)
	fmt.Println(user.Name, ":", message.Content)
	user2 := packet.User{Nickname: user.Name, LoginTime: user.Time.Unix()}
	connect.Broadcast(packet.GlobalMessage{Message: message, User: user2})
}
