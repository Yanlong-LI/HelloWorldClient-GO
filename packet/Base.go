package packet

import (
	// 网关
	_ "github.com/yanlong-li/HelloWorldServer/packet/gateway"
	// 服务器
	_ "github.com/yanlong-li/HelloWorldServer/packet/server"
	// 用户
	_ "github.com/yanlong-li/HelloWorldServer/packet/user"
	// 语音
	_ "github.com/yanlong-li/HelloWorldServer/packet/voice"
)
