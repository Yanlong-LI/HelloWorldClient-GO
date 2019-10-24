package packet

import "HelloWorld/io/network/packet"

func init() {
	packet.Register(0, Connected{})
	packet.Register(1, Disconnect{})
}

// 连接动作
type Connected struct {
}

// 断开连接动作
type Disconnect struct {
}
