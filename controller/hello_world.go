package controller

import (
	"HelloWorld/io/network/packet"
	"HelloWorld/io/network/socket/connect"
	"fmt"
)

func HelloWorld(world packet.HelloWorld, conn *connect.Connector) {

	fmt.Println(world.Message)
	model := packet.HelloWorld{
		Message: "Hello World 我是服务端",
	}
	conn.Send(model)
}

func WriteUint16(n uint16) []byte {
	return []byte{byte(n), byte(n >> 8)}
}
