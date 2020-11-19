package packet_model

import (
	"github.com/yanlong-li/hi-go-gateway/common"
	"github.com/yanlong-li/hi-go-socket/packet"
)

func init() {
	packet.Register(common.ServerAndClientGroup, packet.Connection,
		Connected{},
		Disconnect{},
		BeforeReceiving{},
		BeforeSending{},
	)
}

// 连接动作
type Connected struct {
}

// 断开连接动作
type Disconnect struct {
}

// 接收前置
type BeforeReceiving struct {
}

// 发送前置
type BeforeSending struct {
}
