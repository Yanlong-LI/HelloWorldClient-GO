package gateway

import "github.com/yanlong-li/HelloWorld-GO/io/network/packet"

func init() {
	packet.Register(7006, Reconnect{})
}

type Reconnect struct {
	Disconnect
	// 重连时间
	ReTime uint64
}
