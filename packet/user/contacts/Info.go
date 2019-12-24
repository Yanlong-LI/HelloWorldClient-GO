package contacts

import (
	"HelloWorld/io/network/packet"
	"HelloWorldServer/packet/trait"
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
	Remarks  string
}

type GetInfoFail struct {
	trait.Fail
}
