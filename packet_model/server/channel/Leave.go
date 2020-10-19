package channel

import (
	"github.com/yanlong-li/hi-go-server/packet_model/trait"
	"github.com/yanlong-li/hi-go-socket/packet"
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
