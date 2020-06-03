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

// 获取联系人列表
type GetList struct {
}

//返回联系人列表
//type List struct {
//	List []Info
//}
type List []Info

// 获取联系人列表失败
type GetListFail struct {
	trait.Fail
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
