package controller

import (
	"HelloWorld/io/network/packet"
	"HelloWorld/io/network/route"
	"HelloWorldServer/data/user"
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
