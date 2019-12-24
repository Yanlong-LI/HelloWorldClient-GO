package chat

import (
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/route"
	conn2 "HelloWorldServer/model/online"
	"HelloWorldServer/packet/server/channel/room/message"
	"log"
	"time"
)

func init() {
	route.Register(message.SendTextMessage{}, TextMessage)
}

func TextMessage(msg message.SendTextMessage, conn connect.Connector) {

	_user, err := conn2.Auth(conn.GetId())
	if err != nil {

		log.Print("收到用户消息：获取用户错误")
		return
	}
	_msg := message.TextMessage{
		SendTextMessage: msg,
		Time:            uint64(time.Now().Unix()),
		Author: struct {
			Id       uint64
			Nickname string
		}{Id: _user.Id, Nickname: _user.Nickname},
	}

	conn.Send(message.SendTextMessageSuccess{TextMessage: _msg})

	conn.Broadcast(_msg, false)
}
