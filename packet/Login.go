package packet

import "HelloWorld/io/network/packet"

func init() {
	packet.Register(6001, Login{})
	packet.Register(6002, Token{})
	packet.Register(6003, LoginFail{})
	packet.Register(6006, Logout{})
	packet.Register(6009, NewUser{})
}

// login
type Login struct {
	Username string
	Password string
}

type NewUser struct {
	User
}
type Token struct {
	Token string
}

type LoginFail struct {
	Code    int32
	Message string
}

type Logout struct {
	User
}
