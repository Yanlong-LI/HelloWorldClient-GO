package voice

import "github.com/yanlong-li/hi-go-socket/packet"

func init() {
	packet.Register(11001, T11001{})
	packet.Register(11004, T11004{})
}

type T11001 struct {
	ChannelId uint64
	Content   string
}

type T11004 struct {
	T11001
}
