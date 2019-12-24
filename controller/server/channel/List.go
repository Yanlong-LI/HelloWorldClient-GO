package channel

import (
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/route"
	"HelloWorldServer/packet/server/channel"
	"strconv"
)

func init() {
	route.Register(channel.GetList{}, actionGetChannelList)
}

func actionGetChannelList(_ channel.GetList, conn connect.Connector) {

	list := channel.List{}
	info := channel.Info{
		Id:         1,
		Name:       "内测一",
		Verify:     true,
		CreateUser: "",
		OwnerUser:  "",
		CreateTime: 1321,
		Public:     true,
		Icon:       "el-icon-s-platform",
	}

	for i := 2; i < 6; i++ {
		_channelInfo := channel.Info{
			Id:         uint64(i),
			Name:       "内测频道" + strconv.Itoa(i),
			Verify:     true,
			CreateUser: "0",
			OwnerUser:  "0",
			CreateTime: 1321,
			Public:     true,
			Icon:       "el-icon-s-platform",
		}
		info.Channels = append(info.Channels, _channelInfo)
	}

	list.List = append(list.List, info)

	conn.Send(list)
}
