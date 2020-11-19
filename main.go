package main

import (
	"fmt"
	"github.com/yanlong-li/hi-go-gateway/common"
	logger "github.com/yanlong-li/hi-go-logger"
	_ "github.com/yanlong-li/hi-go-server/controller"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/socket"
	"time"
)

func main() {

	logger.SetLevel(logger.ALL)
	// 连接网关服务器
	go socket.Client(common.GatewayAndServerGroup, "127.0.0.1:3000")

	// 接受客户端连接
	go socket.Server(common.ServerAndClientGroup, ":3002")
	printCount()
}

func printCount() {

	for {
		fmt.Println(connect.Count())
		time.Sleep(time.Second * 5)
	}
}
