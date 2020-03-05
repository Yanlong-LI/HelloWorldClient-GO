package member

import (
	"github.com/yanlong-li/HelloWorldServer/packetModel/server/channel/role"
)

type Info struct {
	OpenId     string
	Nickname   string
	Roles      []role.Info
	JoinedTime uint64
}
