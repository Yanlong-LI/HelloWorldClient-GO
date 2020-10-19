package message

import (
	"github.com/yanlong-li/hi-go-server/packet_model/trait"
	"github.com/yanlong-li/hi-go-socket/packet"
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
	ServerId     uint64
	ChannelId    uint64
	RandomString string
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
