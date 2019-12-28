package voice

import (
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/route"
	"HelloWorldServer/packet/voice"
)

func init() {
	route.Register(voice.T11001{}, actionTestVoice)
}

func actionTestVoice(_vc voice.T11001, conn connect.Connector) {
	conn.Send(voice.T11004{T11001: _vc})
}
