package channel

import (
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorldServer/packetModel/server/channel"
	"github.com/yanlong-li/HelloWorldServer/packetModel/trait"
)

func init() {
	route.Register(channel.InvitationLinkJoin{}, actionInvitationLinkJoin)
}

func actionInvitationLinkJoin(link channel.InvitationLinkJoin, conn connect.Connector) {

	conn.Send(channel.InvitationLinkJoinFail{Fail: trait.Fail{Message: "邀请链接功能暂未开放"}})
}
