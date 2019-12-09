package packet

import (
	// 网关
	_ "HelloWorldServer/packet/gateway"
	// 服务器
	_ "HelloWorldServer/packet/server"
	// 用户
	_ "HelloWorldServer/packet/user"
)
