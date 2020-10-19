package gateway

import "github.com/yanlong-li/hi-go-socket/packet"

func init() {
	packet.Register(7006, Reconnect{})
}

type Reconnect struct {
	Disconnect
	// 重连时间
	ReTime uint64
}
