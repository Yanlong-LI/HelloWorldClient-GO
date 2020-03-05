package gateway

import (
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorldServer/packetModel/gateway"
)

func init() {
	route.Register(gateway.GetInfo{}, GetInfo)
}

func GetInfo(info gateway.GetInfo, conn connect.Connector) {

	conn.Send(gateway.Info{Name: "Master Service", Version: "1.0.0", Region: "China", CreateTime: 1575703496})
}
