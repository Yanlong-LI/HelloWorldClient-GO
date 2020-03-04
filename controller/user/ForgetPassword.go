package user

import (
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorldServer/packet/trait"
	"github.com/yanlong-li/HelloWorldServer/packet/user/forgetPassword"
)

func init() {
	route.Register(forgetPassword.ForgetPassword{}, actionForgetPassword)
}

func actionForgetPassword(_fp forgetPassword.ForgetPassword, conn connect.Connector) {

	conn.Send(forgetPassword.Fail{Fail: trait.Fail{Message: "当前未开通"}})

}
