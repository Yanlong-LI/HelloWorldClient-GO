package packetModel

import "github.com/yanlong-li/HelloWorld-GO/io/network/packet"

func init() {
	packet.Register(packet.CONNECTION, Connected{})
	packet.Register(packet.DISCONNECTION, Disconnect{})
	packet.Register(packet.BEFORE_RECVING, RecvPacketMiddleware{})
	// 加密的数据
	packet.Register(packet.BEFORE_SENDING, Encrypt{})
	//未加密的数据
	packet.Register(6, Decrypt{})
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

//加密后的数据 1
type Encrypt struct {
	data []byte
}

//未加密的数据 0
type Decrypt struct {
	data []byte
}
