package gateway

import (
	"fmt"
	"github.com/yanlong-li/hi-go-gateway/common"
	"github.com/yanlong-li/hi-go-gateway/packet_model"
	"github.com/yanlong-li/hi-go-gateway/packet_model/server"
	logger "github.com/yanlong-li/hi-go-logger"
	common2 "github.com/yanlong-li/hi-go-server/common"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
	"time"
)

func init() {
	route.Register(common.GatewayAndServerGroup, packet_model.Connected{}, Connected)
	route.Register(common.GatewayAndServerGroup, server.RegisterSuccess{}, func(success server.RegisterSuccess, conn connect.Connector) {
		fmt.Println("网关注册成功！")

		for {
			time.Sleep(time.Second * 5)
			err := conn.Send(server.LoadReport{CurrentLoad: connect.Count()})
			if err != nil {
				break
			}
		}
	})
}

func Connected(conn connect.Connector) {
	logger.Debug("连接到网关", 0, conn.GetId())

	common2.Client = conn

	_ = conn.Send(
		server.Register{
			Name:        "s1",
			Version:     "1.0.0",
			IP:          common2.ServerListenServer.IP,
			Port:        common2.ServerListenServer.Port,
			PeakLoad:    50000,
			OptimumLoad: 10000,
			Weight:      100,
		})
}
