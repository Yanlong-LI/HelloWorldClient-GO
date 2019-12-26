package forgetPassword

import (
	"HelloWorld/io/network/packet"
	"HelloWorldServer/packet/trait"
)

type ForgetPassword struct {
	Email string
}

type Success struct {
}

type Fail struct {
	trait.Fail
}

func init() {
	packet.Register(6013, ForgetPassword{})
	packet.Register(6014, Success{})
	packet.Register(6015, Fail{})
}
