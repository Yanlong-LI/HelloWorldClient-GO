package packetModel

import "github.com/yanlong-li/HelloWorld-GO/io/network/packet"

func init() {
	packet.Register(packet.CONNECTION, Connected{})
	packet.Register(packet.DISCONNECTION, Disconnect{})
	packet.Register(packet.BEFORE_RECVING, BeforeRecving{})
	packet.Register(packet.BEFORE_SENDING, BeforeSending{})
}

// 连接动作
type Connected struct {
}

// 断开连接动作
type Disconnect struct {
}

// 接收前置
type BeforeRecving struct {
}

// 发送前置
type BeforeSending struct {
}
