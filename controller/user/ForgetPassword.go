package user

import (
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/route"
	"HelloWorldServer/packet/trait"
	"HelloWorldServer/packet/user/forgetPassword"
)

func init() {
	route.Register(forgetPassword.ForgetPassword{}, actionForgetPassword)
}

func actionForgetPassword(_fp forgetPassword.ForgetPassword, conn connect.Connector) {

	conn.Send(forgetPassword.Fail{Fail: trait.Fail{Message: "当前未开通"}})

}
