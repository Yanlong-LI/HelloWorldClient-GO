package channel

import (
	"github.com/yanlong-li/hi-go-server/common"
	"github.com/yanlong-li/hi-go-server/model"
	"github.com/yanlong-li/hi-go-server/packet_model/server/channel"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
)

func init() {
	route.Register(channel.GetList{}, actionGetChannelList)
}

func actionGetChannelList(_ channel.GetList, conn connect.Connector) {

	user, _ := common.Auth(conn.GetId())

	// model
	_list := model.GetUserChannels(user.Id)
	// packet
	list := channel.List{}
	for _, cu := range _list {

		_cu, ok := cu.(model.ChannelUser)

		if ok != true {
			logger.Fatal("断言错误", 0, cu)
			continue
		}

		cha, err := _cu.Channel()
		if !err.Status() {
			continue
		}
		createUser, _ := model.GetUserById(cha.CreateUserId)
		ownerUser, _ := model.GetUserById(cha.OwnerUserId)
		info := channel.Info{
			Id:     cha.Id,
			Name:   cha.Name,
			Verify: true,
			CreateUser: struct {
				Id       uint64
				Nickname string
			}{Id: createUser.Id, Nickname: createUser.Nickname},
			OwnerUser: struct {
				Id       uint64
				Nickname string
			}{Id: ownerUser.Id, Nickname: ownerUser.Nickname},
			CreateTime: cha.CreateTime,
			Public:     true,
			Avatar:     cha.Avatar,
			Describe:   cha.Describe,
			Channels:   []channel.Info{},
		}
		ChannelChildrenS := cha.GetChildren()
		for _, channelChildren := range ChannelChildrenS {
			if _channelChildren, ok := channelChildren.(model.Channel); ok {
				_createUser, _ := model.GetUserById(cha.CreateUserId)
				_ownerUser, _ := model.GetUserById(cha.OwnerUserId)
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
					Avatar:     _channelChildren.Avatar,
					Describe:   _channelChildren.Describe,
					Channels:   []channel.Info{},
				}
				info.Channels = append(info.Channels, _channelInfo)
			}

		}

		list = append(list, info)
	}

	_ = conn.Send(list)
}
