package user

import (
	"github.com/yanlong-li/hi-go-server/common"
	"github.com/yanlong-li/hi-go-server/model"
	"github.com/yanlong-li/hi-go-server/packet_model/trait"
	"github.com/yanlong-li/hi-go-server/packet_model/user/register"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
	"strings"
	"time"
)

func init() {
	route.Register(register.Register{}, actionRegisterAccount)
}

func actionRegisterAccount(_register register.Register, conn connect.Connector) {

	// 判断邮箱是否正确
	if !common.VerifyEmailFormat(_register.Account) {
		_ = conn.Send(register.Fail{Fail: trait.Fail{Message: "邮箱格式不正确"}})
		return
	}
	_register.Password = strings.Trim(_register.Password, " 	")

	if len(_register.Password) < 6 {
		_ = conn.Send(register.Fail{Fail: trait.Fail{Message: "密码不合格"}})
		return
	}

	_account := &model.UserAccount{}
	_exists := db.Model(_account).Find().Where(map[interface{}]interface{}{"account": _register.Account, "type": 0}).Exists()
	if _exists {
		_ = conn.Send(register.Fail{Fail: trait.Fail{Message: "邮箱已存在"}})
		return
	}

	_time := uint64(time.Now().Unix())
	_user := &model.User{Nickname: _register.Nickname, CreateTime: _time, UpdateTime: _time, Region: "China", Avatar: "", Language: "zh-chs"}
	err := db.Model(_user).Insert().Insert()
	if err.Empty() {
		_ = conn.Send(register.Fail{Fail: trait.Fail{Message: err.Error()}})
		return
	}
	_account.Account = _register.Account
	_account.UserId = _user.Id
	_account.Type = 0
	_account.CreateTime = _time
	_account.UpdateTime = _time
	err = db.Model(_account).Insert().Insert()
	if err.Empty() {
		_ = conn.Send(register.Fail{Fail: trait.Fail{Message: err.Error()}})
		return
	}
	_password := &model.UserPassword{}
	_password.Type = 0
	_password.UserId = _user.Id
	_password.Password = _register.Password
	_password.CreateTime = _time
	_password.UpdateTime = _time
	err = db.Model(_password).Insert().Insert()
	if err.Empty() {
		_ = conn.Send(register.Fail{Fail: trait.Fail{Message: err.Error()}})
		return
	}
	_ = conn.Send(register.Success{})

}
