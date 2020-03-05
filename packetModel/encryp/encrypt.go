package encryp

import "github.com/yanlong-li/HelloWorld-GO/io/network/packet"

func init() {
	packet.Register(7016, EncryptData{})
}

type EncryptData struct {
	Data []byte
}
