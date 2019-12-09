package gateway

import "HelloWorld/io/network/packet"

func init() {
	packet.Register(7006, Reconnect{})
}

type Reconnect struct {
	Disconnect
	// 重连时间
	ReTime uint64
}
