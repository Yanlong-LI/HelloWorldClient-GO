package data

import (
	"HelloWorld/io/network/socket/connect"
	"time"
)

type User struct {
	Name string
	Time time.Time
}

var List = make(map[*connect.Connector]User)

func Register(conn *connect.Connector, user User) {
	List[conn] = user
}
func GetUser(conn *connect.Connector) User {
	return List[conn]
}

func GetUsers() map[*connect.Connector]User {
	return List
}
