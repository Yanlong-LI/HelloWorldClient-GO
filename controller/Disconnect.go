package controller

import (
	"HelloWorld/io/network/route"
	"HelloWorldServer/data/user"
	"HelloWorldServer/packet"
	"fmt"
)

func init() {
	route.Register(packet.Disconnect{}, Disconnect)
}

func Disconnect(disconnect packet.Disconnect, ID uint32) {

	u := user.GetUser(ID)
	fmt.Println("用户注销：", u.Name)
	user.RemoveUser(ID)

}
