package gateway

import "HelloWorld/io/network/packet"

func init() {
	packet.Register(7003, GetInfo{})
	packet.Register(7004, Info{})
}

type Info struct {
	// 网关名称
	Name string
	// 区域
	Region string
	// 创建时间
	CreateTime uint64
	// 版本
	Version string
}

type GetInfo struct {
}
