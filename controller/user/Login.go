package user

import (
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/route"
	"HelloWorldServer/packet"
	"fmt"
)

func init() {
	route.Register(packet.Login{}, Login)
}

func Login(login packet.Login, conn connect.Connector) {
	fmt.Println(login, conn)
	conn.Send(packet.Ticket{Ticket: "213213313213213"})
}
