package member

import (
	"github.com/yanlong-li/hi-go-server/packet_model/server/channel/role"
)

type Info struct {
	OpenId     string
	Nickname   string
	Roles      []role.Info
	JoinedTime uint64
}
