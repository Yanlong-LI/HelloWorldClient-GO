package Login

import (
	"HelloWorld/io/network/packet"
	"HelloWorldServer/packet/trait"
)

func init() {
	packet.Register(6001, ForEmail{})
	packet.Register(6002, Success{})
	packet.Register(6003, Ticket{})
	packet.Register(6004, AuthTicket{})
	packet.Register(6005, Fail{})
	packet.Register(6006, TicketAuthFail{})
}

//登陆
type ForEmail struct {
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
type Success struct {
	Token string
}

type Fail struct {
	trait.Fail
}

type TicketAuthFail struct {
	Code    uint32
	Message string
}
