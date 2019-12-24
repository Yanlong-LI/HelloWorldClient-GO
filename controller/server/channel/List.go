package channel

import (
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/route"
	"HelloWorldServer/model"
	"HelloWorldServer/packet/server/channel"
)

func init() {
	route.Register(channel.GetList{}, actionGetChannelList)
}

func actionGetChannelList(_ channel.GetList, conn connect.Connector) {

	_list := model.GetChannels()
	list := channel.List{}
	for _, cha := range _list {
		if _cha, ok := cha.(model.Channel); ok {
			createUser, _ := model.GetUserById(_cha.CreateUserId)
			ownerUser, _ := model.GetUserById(_cha.CreateUserId)
			info := channel.Info{
				Id:     _cha.Id,
				Name:   _cha.Name,
				Verify: true,
				CreateUser: struct {
					Id       uint64
					Nickname string
				}{Id: createUser.Id, Nickname: createUser.Nickname},
				OwnerUser: struct {
					Id       uint64
					Nickname string
				}{Id: ownerUser.Id, Nickname: ownerUser.Nickname},
				CreateTime: _cha.CreateTime,
				Public:     true,
				Icon:       _cha.Icon,
			}
			ChannelChildrenS := _cha.GetChildren()
			for _, channelChildren := range ChannelChildrenS {
				if _channelChildren, ok := channelChildren.(model.Channel); ok {
					_createUser, _ := model.GetUserById(_cha.CreateUserId)
					_ownerUser, _ := model.GetUserById(_cha.CreateUserId)
					_channelInfo := channel.Info{
						Id:     _channelChildren.Id,
						Name:   _channelChildren.Name,
						Verify: true,
						CreateUser: struct {
							Id       uint64
							Nickname string
						}{Id: _createUser.Id, Nickname: _createUser.Nickname},
						OwnerUser: struct {
							Id       uint64
							Nickname string
						}{Id: _ownerUser.Id, Nickname: _ownerUser.Nickname},
						CreateTime: _channelChildren.CreateTime,
						Public:     true,
						Icon:       _channelChildren.Icon,
					}
					info.Channels = append(info.Channels, _channelInfo)
				}

			}

			list.List = append(list.List, info)
		}
	}

	conn.Send(list)
}
