package gateway

import "github.com/yanlong-li/HelloWorld-GO/io/network/packet"

func init() {
	packet.Register(7003, Config{})
}

// 网关配置 7001
type Config struct {
	// 心跳间隔 毫秒级别
	HeartbeatInterval uint32
	Info
}
