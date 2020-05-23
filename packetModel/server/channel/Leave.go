package channel

import (
	"github.com/yanlong-li/HelloWorld-GO/io/network/packet"
	"github.com/yanlong-li/HelloWorldServer/packetModel/trait"
)

// 离开频道
type LeaveChannel struct {
	Id uint64
}
type LeaveChannelSuccess struct {
	Id uint64
}
type LeaveChannelFail struct {
	trait.Fail
}

func init() {
	packet.Register(9079, LeaveChannel{}, LeaveChannelSuccess{}, LeaveChannelFail{})
}
