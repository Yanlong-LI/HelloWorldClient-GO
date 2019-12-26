package channel

import (
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/route"
	"HelloWorldServer/packet/server/channel"
	"HelloWorldServer/packet/trait"
)

func init() {
	route.Register(channel.InvitationLinkJoin{}, actionInvitationLinkJoin)
}

func actionInvitationLinkJoin(link channel.InvitationLinkJoin, conn connect.Connector) {

	conn.Send(channel.InvitationLinkJoinFail{Fail: trait.Fail{Message: "邀请链接功能暂未开放"}})
}
