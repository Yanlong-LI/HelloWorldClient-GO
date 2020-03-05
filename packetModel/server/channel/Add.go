package channel

import (
	"github.com/yanlong-li/HelloWorld-GO/io/network/packet"
	"github.com/yanlong-li/HelloWorldServer/packetModel/trait"
)

type AddChannel struct {
	Name string
	//Region   string
	Avatar   string
	Describe string
}
type AddChannelSuccess struct {
	Info
}
type AddChannelFail struct {
	trait.Fail
}

type AddSubChannel struct {
	Name      string
	Avatar    string
	Describe  string
	ChannelId uint64
}

type AddSubChannelSuccess struct {
	ChannelId uint64
	Channel   Info
}
type AddSubChannelFail struct {
	trait.Fail
}

func init() {
	packet.Register(9019, AddChannel{})
	packet.Register(9020, AddChannelSuccess{})
	packet.Register(9021, AddChannelFail{})
	packet.Register(9085, AddSubChannel{})
	packet.Register(9086, AddSubChannelSuccess{})
	packet.Register(9087, AddSubChannelFail{})
}
