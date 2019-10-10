package packet

import "HelloWorld/io/network/packet"

func init() {
	packet.Register(6004, GetUserList{})
	packet.Register(6005, UserList{List: make([]string, 0)})
	packet.Register(1, Disconnect{})
}

type GetUserList struct {
}

type UserList struct {
	List []string
}
type Disconnect struct {
}
