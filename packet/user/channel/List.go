package channel

import (
	"github.com/yanlong-li/HelloWorld-GO/io/network/packet"
	"github.com/yanlong-li/HelloWorldServer/packet/server/channel"
	"github.com/yanlong-li/HelloWorldServer/packet/trait"
)

func init() {
	packet.Register(6513, GetJoinChannelList{})
	packet.Register(6514, JoinChannelList{})
	packet.Register(6515, GetJoinChannelListFail{})
}

type GetJoinChannelList struct {
}
type JoinChannelList struct {
	List []channel.Info
}
type GetJoinChannelListFail struct {
	trait.Fail
}
