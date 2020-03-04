package message

import (
	"github.com/yanlong-li/HelloWorld-GO/io/network/packet"
	"github.com/yanlong-li/HelloWorldServer/packet/trait"
)

func init() {
	packet.Register(6601, SendTextMessage{})
	packet.Register(6602, SendMessageSuccess{})
	packet.Register(6603, SendMessageFail{})
	packet.Register(6604, RecvTextMessage{})
}

type SendTextMessage struct {
	ContactId    uint64
	Content      string
	RandomString string
}

type SendMessageSuccess struct {
	Id uint64
	SendTextMessage
	CreateTime uint64
}
type SendMessageFail struct {
	trait.Fail
}

type RecvTextMessage struct {
	Id         uint64
	ContactId  uint64
	Content    string
	CreateTime uint64
}
