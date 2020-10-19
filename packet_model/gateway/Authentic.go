package gateway

import (
	"github.com/yanlong-li/hi-go-server/packet_model/trait"
	"github.com/yanlong-li/hi-go-socket/packet"
)

func init() {
	packet.Register(7015, AuthenticFail{})
}

type AuthenticFail struct {
	trait.Fail
}
