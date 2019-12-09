package channel

import (
	"HelloWorld/io/network/packet"
	"HelloWorldServer/packet/server/channel"
	"HelloWorldServer/packet/trait"
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
