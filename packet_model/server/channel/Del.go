package channel

import (
	"github.com/yanlong-li/hi-go-server/packet_model/trait"
	"github.com/yanlong-li/hi-go-socket/packet"
)

type DelChannel struct {
	Id uint64
}
type DelChannelSuccess struct {
	trait.Success
	Id uint64
}
type DelChannelFail struct {
	trait.Fail
}

func init() {
	packet.Register(9016, DelChannel{})
	packet.Register(9017, DelChannelSuccess{})
	packet.Register(9018, DelChannelFail{})
}
