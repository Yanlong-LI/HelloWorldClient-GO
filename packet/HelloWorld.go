package packet

import "HelloWorld/io/network/packet"

func init() {
	packet.Register(0x00, HelloWorld{})
}

// hello world
type HelloWorld struct {
	Message string
}
