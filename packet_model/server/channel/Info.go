package channel

import (
	"github.com/yanlong-li/hi-go-server/packet_model/trait"
	"github.com/yanlong-li/hi-go-socket/packet"
)

func init() {
	packet.Register(9010, Get{}, Info{}, GetFail{})
}

type Get struct {
	Id uint64
}

type Info struct {
	Id         uint64
	Name       string
	Avatar     string
	CreateUser struct {
		Id       uint64
		Nickname string
	}
	OwnerUser struct {
		Id       uint64
		Nickname string
	} // 实际掌控着
	CreateTime uint64
	Public     bool //是否公开
	Verify     bool // 是否经过验证
	Commerce   bool // 是否可商业
	Channels   []Info
	Describe   string
}

type GetFail struct {
	trait.Fail
}
