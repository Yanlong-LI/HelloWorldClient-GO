package packet

import "HelloWorld/io/network/packet"

func init() {
	packet.Register(6007, Message{})
	packet.Register(6008, GlobalMessage{})
}

// 全局群发消息
type Message struct {
	Content string
}

// 全局消息发送包装
type GlobalMessage struct {
	Message
	User
}
