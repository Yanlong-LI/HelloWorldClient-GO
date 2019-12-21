package channel

import "HelloWorld/io/network/packet"

func init() {
	packet.Register(9010, Get{})
	packet.Register(9011, Info{})
}

type Get struct {
	Id uint64
}

type Info struct {
	Id         uint64
	Name       string
	Icon       string
	CreateUser string
	OwnerUser  string // 实际掌控着
	CreateTime uint64
	Public     bool //是否公开
	Verify     bool // 是否经过验证
	Commerce   bool // 是否可商业
	Channels   []Info
}
