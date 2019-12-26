package register

import (
	"HelloWorld/io/network/packet"
	"HelloWorldServer/packet/trait"
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
