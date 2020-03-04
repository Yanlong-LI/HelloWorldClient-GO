package gateway

import (
	"github.com/yanlong-li/HelloWorld-GO/io/network/packet"
	"github.com/yanlong-li/HelloWorldServer/packet/trait"
)

func init() {
	packet.Register(7015, AuthenticFail{})
}

type AuthenticFail struct {
	trait.Fail
}
