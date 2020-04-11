package controller

import (
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorld-GO/io/network/stream"
	conn2 "github.com/yanlong-li/HelloWorldServer/model/online"
	"github.com/yanlong-li/HelloWorldServer/packetModel"
	"github.com/yanlong-li/HelloWorldServer/packetModel/gateway"
	"github.com/yanlong-li/HelloWorldServer/packetModel/trait"
)

// 路由白名单
var routingWhiteList = make(map[uint32]bool, 0)

// 这里实现的是 账户登录白名单，可以自由定义其他类型检查
func init() {
	route.Register(packetModel.BeforeRecving{}, Middleware)

	routingWhiteList[7001] = true // 心跳
	routingWhiteList[7002] = true // 获取心跳
	routingWhiteList[6001] = true // 邮箱登录
	routingWhiteList[6007] = true // 恢复数据包
	routingWhiteList[6010] = true // 账户注册
	routingWhiteList[6013] = true // 忘记密码
	//加密数据放行
	routingWhiteList[7016] = true
}

func Middleware(ps stream.Interface, conn connect.Connector) bool {
	if _, ok := routingWhiteList[ps.GetOpCode()]; !ok {
		// 验证用户是否登陆
		_, err := conn2.Auth(conn.GetId())
		if err != nil {
			_ = conn.Send(gateway.AuthenticFail{Fail: trait.Fail{Code: 7015, Message: "当前未登陆"}})
			return false
		}

	}

	return true
}
