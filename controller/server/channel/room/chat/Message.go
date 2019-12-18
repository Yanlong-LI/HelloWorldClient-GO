package chat

import (
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/route"
	conn2 "HelloWorldServer/model/Login"
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
			Id       string
			Username string
		}{Id: _user.Id, Username: _user.UserName},
	}

	conn.Send(message.SendTextMessageSuccess{TextMessage: _msg})

	if msg.ChannelId == "@me" {
		//todo 单发给某个用户
		//conn.Send(_msg)
	} else {
		//todo 群发给已加入群组的用户
		conn.Broadcast(_msg, false)
	}
}
