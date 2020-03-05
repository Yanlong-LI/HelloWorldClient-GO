package controller

import (
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorld-GO/io/network/socket/stream"
	"github.com/yanlong-li/HelloWorldServer/packetModel"
	"github.com/yanlong-li/HelloWorldServer/packetModel/encryp"
)

//加密和解密

func init() {
	route.Register(packetModel.BeforeSending{}, Encryption)
}

// 加密动作
func Encryption(OpCode uint32, Data []byte) []byte {

	//简单做了一层封装
	model := encryp.EncryptData{}
	model.Data = Data

	ps := stream.SocketPacketStream{}
	ps.Marshal(model)

	return ps.ToData()
}
