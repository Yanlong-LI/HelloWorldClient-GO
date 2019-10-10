package controller

import (
	"HelloWorld/io/network/route"
	"HelloWorld/io/network/socket/connect"
	user2 "HelloWorldServer/data/user"
	"HelloWorldServer/packet"
	"fmt"
)

func init() {
	route.Register(packet.GetUserList{}, GetUserList)
}

func GetUserList(list packet.GetUserList, conn *connect.Connector) {
	user := user2.GetUser(conn.ID)

	var userList []string
	for _, v := range user2.GetUsers() {
		userList = append(userList, v.Name)
	}

	fmt.Println(user)
	conn.Send(packet.UserList{List: userList})
}
