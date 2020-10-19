package voice

import (
	"github.com/yanlong-li/hi-go-server/packet_model/voice"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
)

func init() {
	route.Register(voice.T11001{}, actionTestVoice)
}

func actionTestVoice(_vc voice.T11001, conn connect.Connector) {
	_ = conn.Send(voice.T11004{T11001: _vc})
}
