package message

import (
	"github.com/yanlong-li/HelloWorld-GO/io/network/packet"
	"github.com/yanlong-li/HelloWorldServer/packetModel/trait"
)

func init() {
	packet.Register(10001, SendTextMessage{})
	packet.Register(10002, SendTextMessageSuccess{})
	packet.Register(10003, SendTextMessageFail{})
	packet.Register(10004, TextMessage{})
}

type SendTextMessage struct {
	ServerId     uint64
	ChannelId    uint64
	Content      string
	RandomString string
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
	Id     uint64
	Time   uint64
	Author struct {
		Id       uint64
		Nickname string
	}
}
