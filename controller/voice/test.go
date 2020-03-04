package voice

import (
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorldServer/packet/voice"
)

func init() {
	route.Register(voice.T11001{}, actionTestVoice)
}

func actionTestVoice(_vc voice.T11001, conn connect.Connector) {
	conn.Send(voice.T11004{T11001: _vc})
}
