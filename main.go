package main

import (
	"fmt"
	"github.com/yanlong-li/hi-go-gateway/common"
	logger "github.com/yanlong-li/hi-go-logger"
	common2 "github.com/yanlong-li/hi-go-server/common"
	_ "github.com/yanlong-li/hi-go-server/controller"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/socket"
	"net"
	"strings"
	"time"
)

func main() {
	logger.SetLevel(logger.INFO)

	address, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		return
	}
	ip := address[1]
	index := strings.Index(ip.String(), "/")

	server := common2.Server{
		IP:   net.ParseIP(ip.String()[:index]),
		Port: 3002,
	}
	common2.ServerListenServer = server

	go gateway()

	// 接受客户端连接
	go socket.Server(common.ServerAndClientGroup, ":3002")
	printCount()
}

func gateway() {
	defer func() {
		gateway()
	}()
	// 连接网关服务器
	socket.Client(common.GatewayAndServerGroup, "gateway:3000")
}

func printCount() {

	for {
		fmt.Println(connect.Count())
		time.Sleep(time.Second * 5)
	}
}
