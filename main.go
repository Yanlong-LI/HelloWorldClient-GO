package main

import (
	"HelloWorld/io/network/socket"
	_ "HelloWorldServer/controller"
	_ "HelloWorldServer/packet"
)

func main() {

	socket.Server()
	//websocket.Server()
}
