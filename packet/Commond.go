package packet

import "github.com/yanlong-li/HelloWorld-GO/io/network/packet"

func init() {
	packet.Register(0, Connected{})
	packet.Register(1, Disconnect{})
	packet.Register(2, RecvPacketMiddleware{})
}

// 连接动作
type Connected struct {
}

// 断开连接动作
type Disconnect struct {
}

// 收报中间件
type RecvPacketMiddleware struct {
}
