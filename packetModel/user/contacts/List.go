package contacts

import (
	"github.com/yanlong-li/HelloWorld-GO/io/network/packet"
	"github.com/yanlong-li/HelloWorldServer/packetModel/trait"
)

func init() {
	packet.Register(6507, GetList{})
	packet.Register(6508, List{})
	packet.Register(6509, GetListFail{})

	packet.Register(6540, GetBlacklist{}, Blacklist{}, GetBlacklistFail{})
	packet.Register(6546, GetRequestList{}, RequestList{}, GetRequestListFail{})
}

type GetList struct {
}

type List struct {
	List []Info
}

// 获取新朋友列表
type GetRequestList struct {
}

// 新朋友列表
type RequestList struct {
	List []Info
}

//获取新朋友列表失败
type GetRequestListFail struct {
	trait.Fail
}

type GetListFail struct {
	trait.Fail
}

// 获取 黑名单列表
type GetBlacklist struct {
}

// 黑名单列表
type Blacklist struct {
	List []Info
}
type GetBlacklistFail struct {
	trait.Fail
}
