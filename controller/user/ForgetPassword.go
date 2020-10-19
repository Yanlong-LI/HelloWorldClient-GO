package user

import (
	"github.com/yanlong-li/hi-go-server/packet_model/trait"
	"github.com/yanlong-li/hi-go-server/packet_model/user/forgetPassword"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
)

func init() {
	route.Register(forgetPassword.ForgetPassword{}, actionForgetPassword)
}

func actionForgetPassword(_fp forgetPassword.ForgetPassword, conn connect.Connector) {

	_ = conn.Send(forgetPassword.Fail{Fail: trait.Fail{Message: "当前未开通"}})

}
