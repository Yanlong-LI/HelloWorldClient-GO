package model

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(&User{})
}

type User struct {
	Id       uint64
	Nickname string
}

// 用户账户表
type UserAccount struct {
	Id      uint64
	UserId  uint64
	Account string
	Type    uint8 // 0 邮箱 1 手机号码
}

// 用户密码表
type UserPassword struct {
	Id       uint64
	UserId   uint64
	Type     uint8 // 0登陆密码 1 支付密码
	Password string
}
