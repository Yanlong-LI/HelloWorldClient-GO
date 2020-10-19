package register

import (
	"github.com/yanlong-li/hi-go-server/packet_model/trait"
	"github.com/yanlong-li/hi-go-socket/packet"
)

type Register struct {
	Account  string
	Type     uint8
	Password string
	Nickname string
}

type Success struct {
}

type Fail struct {
	trait.Fail
}

func init() {
	packet.Register(6010, Register{})
	packet.Register(6011, Success{})
	packet.Register(6012, Fail{})
}
