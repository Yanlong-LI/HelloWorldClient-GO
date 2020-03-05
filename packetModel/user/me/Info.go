package me

import (
	"github.com/yanlong-li/HelloWorld-GO/io/network/packet"
	"github.com/yanlong-li/HelloWorldServer/packetModel/trait"
)

func init() {
	packet.Register(6501, GetInfo{})
	packet.Register(6502, Info{})
	packet.Register(6503, GetInfoFail{})
	packet.Register(6004, EditInfo{})
	packet.Register(6005, EditSuccess{})
	packet.Register(6006, EditFail{})
}

//region 获取用户信息

type GetInfo struct {
}

type Info struct {
	Id       uint64 //用户ID
	Nickname string //用户昵称
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

type EditInfo struct {
	Nickname string
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
