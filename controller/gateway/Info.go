package gateway

import (
	logger "github.com/yanlong-li/hi-go-logger"
	"github.com/yanlong-li/hi-go-server/packet_model/gateway"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
)

func init() {
	route.Register(gateway.GetInfo{}, GetInfo)
}

func GetInfo(info gateway.GetInfo, conn connect.Connector) {

	logger.Debug("发送网关信息", 0)
	_ = conn.Send(gateway.Info{Name: "Master Service", Version: "1.0.0", Region: "China", CreateTime: 1575703496})

}
