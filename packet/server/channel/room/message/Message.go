package message

import (
	"HelloWorld/io/network/packet"
	"HelloWorldServer/packet/trait"
)

func init() {
	packet.Register(10001, SendTextMessage{})
	packet.Register(10002, SendTextMessageSuccess{})
	packet.Register(10003, SendTextMessageFail{})
	packet.Register(10004, TextMessage{})
}

type SendTextMessage struct {
	ServerId  string
	ChannelId string
	RoomId    string
	Content   string
	RandomStr string
}

type SendTextMessageSuccess struct {
	trait.Success
	TextMessage
}
type SendTextMessageFail struct {
	trait.Fail
}

type TextMessage struct {
	SendTextMessage
	Id     string
	Time   uint64
	Author struct {
		Id       string
		Username string
	}
}
