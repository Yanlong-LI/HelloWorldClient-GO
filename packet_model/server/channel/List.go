package channel

import "github.com/yanlong-li/hi-go-socket/packet"

func init() {
	packet.Register(9001, GetList{})
	packet.Register(9002, List{})
	packet.Register(9004, SearchChannelList{})
	packet.Register(9005, SearchChannelListSuccess{})
}

type GetList struct {
}
type List []Info
type SearchChannelList struct {
	Name string
}
type SearchChannelListSuccess []Info
