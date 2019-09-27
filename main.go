package main

import (
	"HelloWorld/io/network/socket"
	_ "HelloWorldServer/controller"
)

func main() {

	socket.Server()
}
