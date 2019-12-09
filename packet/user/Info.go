package user

import (
	"HelloWorld/io/network/packet"
	"HelloWorldServer/packet/trait"
)

func init() {
	packet.Register(6501, GetInfo{})
	packet.Register(6502, User{})
	packet.Register(6503, GetInfoFail{})
	packet.Register(6004, Edit{})
	packet.Register(6005, EditSuccess{})
	packet.Register(6006, EditFail{})
}

//region 获取用户信息

type GetInfo struct {
}

type User struct {
	Id       string //用户ID
	UserName string //用户昵称
	Avatar   string //用户头像
	//Email       string
	//EmailVerify bool
	Language string
	Region   string
	//RegisterTime uint64 //注册时间
}

type GetInfoFail struct {
	trait.Fail
}

//endregion

//region 修改信息

type Edit struct {
	UserName string
	Avatar   string
	Language string
	Region   string
}

type EditSuccess struct {
	trait.Success
}
type EditFail struct {
	trait.Fail
}

//endregion
