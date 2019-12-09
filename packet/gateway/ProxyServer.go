package gateway

import "HelloWorld/io/network/packet"

func init() {
	packet.Register(7009, GetProxyServerList{})
	packet.Register(7010, ProxyServerList{})
	packet.Register(7012, GetProxyInfo{})
}

type GetProxyServerList struct {
}

type ProxyServerList struct {
	List []ProxyServer
}

type GetProxyInfo struct {
	Id string
}
type ProxyServer struct {
	Id      string
	Name    string
	Host    string
	Version string
	Region  string
	RunTime uint64
}
