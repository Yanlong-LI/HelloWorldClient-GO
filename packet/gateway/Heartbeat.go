package gateway

import "HelloWorld/io/network/packet"

func init() {
	packet.Register(7001, Heartbeat{})
	packet.Register(7002, GetHeartbeat{})
}

type Heartbeat struct {
	Sn uint64
}

type GetHeartbeat struct {
}