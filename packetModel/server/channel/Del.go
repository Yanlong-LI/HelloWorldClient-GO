package channel

import (
	"github.com/yanlong-li/HelloWorld-GO/io/network/packet"
	"github.com/yanlong-li/HelloWorldServer/packetModel/trait"
)

type DelChannel struct {
	Id uint64
}
type DelChannelSuccess struct {
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
