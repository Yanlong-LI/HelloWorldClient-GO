package encrypt

import "github.com/yanlong-li/hi-go-socket/packet"

func init() {
	packet.Register(7016, BytesData{})
}

type BytesData struct {
	Data []byte
}
