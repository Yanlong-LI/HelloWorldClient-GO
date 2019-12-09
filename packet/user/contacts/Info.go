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
	Id string
}

type Info struct {
	Id       string
	UserName string
	Avatar   string
	Language string
	Region   string
}

type GetInfoFail struct {
	trait.Fail
}
