package main

import (
	"HelloWorld/io/network/handle"
	"HelloWorld/io/network/socket"
	"HelloWorldServer/controller"
)

func main() {
	handle.Register(0x0001, controller.HelloWorld)
	socket.Server()
}
