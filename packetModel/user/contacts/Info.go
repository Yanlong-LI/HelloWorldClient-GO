package contacts

import (
	"github.com/yanlong-li/HelloWorld-GO/io/network/packet"
	"github.com/yanlong-li/HelloWorldServer/packetModel/trait"
)

func init() {
	packet.Register(6510, GetInfo{})
	packet.Register(6511, Info{})
	packet.Register(6512, GetInfoFail{})
}

type GetInfo struct {
	Id uint64
}

type Info struct {
	Id       uint64
	Nickname string
	Avatar   string
	Language string
	Region   string
	Remark   string
	Online   bool
}

type GetInfoFail struct {
	trait.Fail
}
