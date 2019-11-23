package packet

import "HelloWorld/io/network/packet"

func init() {
	packet.Register(6001, Login{})
	packet.Register(6002, LoginSuccess{})
	packet.Register(6003, Ticket{})
	packet.Register(6004, AuthTicket{})
}

//登陆
type Login struct {
	//邮箱
	Email string
	//密码
	Password string
}

//登陆票据 【可选】
type Ticket struct {
	Ticket string
}

//登陆票据验证
type AuthTicket struct {
	Ticket
	//票据验证码
	Code string
}

//登陆成功
type LoginSuccess struct {
	Token string
}
