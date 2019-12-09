package gateway

import "HelloWorld/io/network/packet"

func init() {
	packet.Register(7005, Disconnect{})
}

type Disconnect struct {
	// 断开信息
	Message string
	// 断开时间
	Time uint64
}
