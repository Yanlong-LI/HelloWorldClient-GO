package controller

import (
	"HelloWorld/io/network/packet"
	"HelloWorld/io/network/route"
	"HelloWorld/io/network/socket/connect"
	"HelloWorldServer/data"
	"fmt"
)

func init() {
	route.Register(packet.GetUserList{}, GetUserList)
}

func GetUserList(list packet.GetUserList, conn *connect.Connector) {
	user := data.GetUser(conn)

	var userList []string
	for _, v := range data.GetUsers() {
		userList = append(userList, v.Name)
	}

	fmt.Println(user)
	conn.Send(packet.UserList{List: userList})
}
