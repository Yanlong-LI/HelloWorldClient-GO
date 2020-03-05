package controller

import (
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	connect2 "github.com/yanlong-li/HelloWorld-GO/io/network/socket/connect"
	"github.com/yanlong-li/HelloWorldServer/packetModel"
)

//加密和解密

func init() {
	route.Register(packetModel.Encrypt{}, Encryption)
	route.Register(packetModel.Decrypt{}, Decryption)
}

// 加密动作
func Encryption(OpCode uint32, Data []byte) []byte {

	//opcode 用于处理是否要加密和加密方式
	data := make([]byte, 0, len(Data)+4)
	// 加密标识数据包
	data = append(data, connect.WriteUint16(uint16(len(Data)+4))...)
	data = append(data, connect.Uint32ToHex(7016)...)
	// 填充数据包
	data = append(data, Data...)
	return data
}

// 解密动作
func Decryption(data []byte, conn connect2.SocketConnector) {
	conn.HandleData(data)
}
