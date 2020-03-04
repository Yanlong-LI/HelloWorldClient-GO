package channel

import (
	"github.com/yanlong-li/HelloWorld-GO/io/network/packet"
	"github.com/yanlong-li/HelloWorldServer/packet/trait"
)

func init() {
	packet.Register(9022, InvitationLinkJoin{})
	packet.Register(9023, InvitationLinkJoinSuccess{})
	packet.Register(9024, InvitationLinkJoinFail{})
}

// 通过邀请链接加入
type InvitationLinkJoin struct {
	InvitationLink string
}

// 加入成功 发送服务器信息
type InvitationLinkJoinSuccess struct {
	Info
}
type InvitationLinkJoinFail struct {
	trait.Fail
}
