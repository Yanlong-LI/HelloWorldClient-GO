package controller

import (
	"github.com/yanlong-li/hi-go-server/packet_model"
	"github.com/yanlong-li/hi-go-server/packet_model/gateway"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
)

func init() {
	route.Register(packet_model.Connected{}, Connected)
}

func Connected(conn connect.Connector) {

	logger.Debug("新连接", 0, conn.GetId())
	_ = conn.Send(gateway.Info{Name: "Master Service", Version: "1.0.0", Region: "China", CreateTime: 1575703496})
	//conn.Send(gateway.GetHeartbeat{})
}
