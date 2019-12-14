package gateway

import (
	"HelloWorld/io/network/packet"
	"HelloWorldServer/packet/trait"
)

func init() {
	packet.Register(7015, AuthenticFail{})
}

type AuthenticFail struct {
	trait.Fail
}
