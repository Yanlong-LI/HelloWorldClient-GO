package encrypt

import "github.com/yanlong-li/HelloWorld-GO/io/network/packet"

func init() {
	packet.Register(7016, BytesData{})
}

type BytesData struct {
	Data []byte
}
