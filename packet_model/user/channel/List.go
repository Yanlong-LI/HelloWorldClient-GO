package channel

import (
	"github.com/yanlong-li/hi-go-server/packet_model/server/channel"
	"github.com/yanlong-li/hi-go-server/packet_model/trait"
	"github.com/yanlong-li/hi-go-socket/packet"
)

func init() {
	packet.Register(6513, GetJoinChannelList{})
	packet.Register(6514, JoinChannelList{})
	packet.Register(6515, GetJoinChannelListFail{})
}

type GetJoinChannelList struct {
}
type JoinChannelList []channel.Info
type GetJoinChannelListFail struct {
	trait.Fail
}
