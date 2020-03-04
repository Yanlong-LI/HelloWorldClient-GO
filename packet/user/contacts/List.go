package contacts

import (
	"github.com/yanlong-li/HelloWorld-GO/io/network/packet"
	"github.com/yanlong-li/HelloWorldServer/packet/trait"
)

func init() {
	packet.Register(6507, GetList{})
	packet.Register(6508, List{})
	packet.Register(6509, GetListFail{})
}

type GetList struct {
}

type List struct {
	List []Info
}

type GetListFail struct {
	trait.Fail
}
