package channel

import "HelloWorld/io/network/packet"

func init() {
	packet.Register(9001, GetList{})
	packet.Register(9002, List{})
}

type GetList struct {
}
type List struct {
	List []Info
}
