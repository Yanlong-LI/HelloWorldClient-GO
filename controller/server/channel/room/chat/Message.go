package chat

import (
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/route"
	"HelloWorldServer/packet/server/channel/room/message"
	"fmt"
)

func init() {
	route.Register(message.SendTextMessage{}, TextMessage)
}

func TextMessage(msg message.SendTextMessage, conn connect.Connector) {
	fmt.Println("收到消息")
	fmt.Println(msg)
	conn.Send(message.TextMessage{
		SendTextMessage: msg,
		Time:            201900000000,
		Author: struct {
			Id       string
			Username string
		}{
			Id:       "user-2",
			Username: "测试1好",
		},
	})
}
