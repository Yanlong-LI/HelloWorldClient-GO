package controller

import (
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/route"
	"HelloWorldServer/data/user"
	"HelloWorldServer/packet"
	"fmt"
)

func init() {
	route.Register(packet.GetUserList{}, GetUserList)
}

func GetUserList(list packet.GetUserList, conn connect.Connector) {
	userInfo, ok := user.GetUser(conn.GetId())
	if ok {
		var userList []string
		for _, v := range user.GetUsers() {
			userList = append(userList, v.Name)
		}

		fmt.Println(userInfo)
		conn.Send(packet.UserList{List: userList})
	}
}
