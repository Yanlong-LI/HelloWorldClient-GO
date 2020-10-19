package Login

import (
	"github.com/yanlong-li/hi-go-server/packet_model/trait"
	"github.com/yanlong-li/hi-go-socket/packet"
)

func init() {
	packet.Register(6001, ForEmail{})
	packet.Register(6002, Success{})
	packet.Register(6003, Ticket{})
	packet.Register(6004, AuthTicket{})
	packet.Register(6005, Fail{})
	packet.Register(6006, TicketAuthFail{})
	packet.Register(6007, Resuming{})
	packet.Register(6008, ResumingSuccess{})
	packet.Register(6009, ResumingFail{})
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

//失败
type Fail struct {
	trait.Fail
}

// 票据验证失败
type TicketAuthFail struct {
	Code    uint32
	Message string
}

// 恢复数据包
type Resuming struct {
	Token string
}

//恢复成功
type ResumingSuccess struct {
	trait.Success
}

//恢复失败
type ResumingFail struct {
	trait.Fail
}
