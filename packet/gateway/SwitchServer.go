package gateway

import "github.com/yanlong-li/HelloWorld-GO/io/network/packet"

func init() {
	packet.Register(7007, SwitchServer{})
}

type SwitchServer struct {
	// 目标服务器
	Host string
	// 信息
	Message string
	// 服务器信息
	Info Info
}
