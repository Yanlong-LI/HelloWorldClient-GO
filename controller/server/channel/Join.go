package channel

import (
	"github.com/yanlong-li/hi-go-server/packet_model/server/channel"
	"github.com/yanlong-li/hi-go-server/packet_model/trait"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
)

func init() {
	route.Register(channel.InvitationLinkJoin{}, actionInvitationLinkJoin)
}

func actionInvitationLinkJoin(link channel.InvitationLinkJoin, conn connect.Connector) {

	_ = conn.Send(channel.InvitationLinkJoinFail{Fail: trait.Fail{Message: "邀请链接功能暂未开放"}})
}
