package contacts

import (
	"github.com/yanlong-li/HelloWorld-GO/io/network/packet"
	"github.com/yanlong-li/HelloWorldServer/packetModel/trait"
	"github.com/yanlong-li/HelloWorldServer/packetModel/user"
)

func init() {
	packet.Register(6519, SearchUser{}, SearchUserSuccess{}, SearchUserFail{})

	//注册发送添加联系人请求
	packet.Register(6522, AddContact{}, AddContactSuccess{}, AddContactFail{})

	packet.Register(6525, AcceptContact{}, AcceptContactSuccess{}, AcceptContactFail{})

	packet.Register(6528, RefuseContact{}, RefuseContactSuccess{}, RefuseContactFail{})

	packet.Register(6531, DeleteContact{}, DeleteContactSuccess{}, DeleteContactFail{})

	packet.Register(6534, BlackContact{}, BlackContactSuccess{}, BlackContactFail{})

	packet.Register(6537, UnBlackContact{}, UnBlackContactSuccess{}, UnBlackContactFail{})

	packet.Register(6543, AcceptedContact{}, RejectedContactSuccess{}, RequestAddContact{})

}

// 搜索账户
type SearchUser struct {
	Account string //账户
}
type SearchUserFail struct {
	trait.Fail
}
type SearchUserSuccess struct {
	user.Info
}

// 添加联系人
type AddContact struct {
	Id     uint64
	Remark string //说明
}

// 添加联系人请求发送成功
type AddContactSuccess struct {
	trait.Success
}

// 添加联系人失败
type AddContactFail struct {
	trait.Fail
}

//接受
type AcceptContact struct {
	Id uint64
}
type AcceptContactSuccess struct {
	trait.Success
}
type AcceptContactFail struct {
	trait.Fail
}

//拒绝
type RefuseContact struct {
	Id uint64
}
type RefuseContactSuccess struct {
	trait.Success
}
type RefuseContactFail struct {
	trait.Fail
}

// 删除
type DeleteContact struct {
	Id    uint64
	Block bool
}
type DeleteContactSuccess struct {
	trait.Success
}
type DeleteContactFail struct {
	trait.Fail
}

// 拉黑
type BlackContact struct {
	Id uint64
}
type BlackContactSuccess struct {
	trait.Success
}
type BlackContactFail struct {
	trait.Fail
}

// 取消拉黑
type UnBlackContact struct {
	Id uint64
}
type UnBlackContactSuccess struct {
	trait.Success
}
type UnBlackContactFail struct {
	trait.Fail
}

//被接受
type AcceptedContact struct {
	Id uint64 // 谁接受的发送谁的id
}

// 被拒绝
type RejectedContactSuccess struct {
	Id     uint64 // 谁拒绝的发送谁的id
	Remark string // 拒绝理由
}

// 添加联系人
type RequestAddContact struct {
	AddContact
}

// 被忽略
type IgnoredContactSuccess struct {
	Id uint64 // 谁忽略的发送谁的id
}
