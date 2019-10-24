package main

import (
	"HelloWorld/io/network/socket"
	"HelloWorld/io/network/websocket"
	_ "HelloWorldServer/controller"
	_ "HelloWorldServer/packet"
)

func main() {

	go websocket.Server(":3001")
	socket.Server(":3000")
}
