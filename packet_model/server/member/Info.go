package member

import "github.com/yanlong-li/hi-go-server/packet_model/server/role"

// 想给某个用户设置某种角色，当然需要明确给谁设定了
type Info struct {
	//用户id
	UserId  uint64
	UnionId string
	Roles   []role.Info
}
