package channel

import (
	db "github.com/yanlong-li/hi-go-orm"
	"github.com/yanlong-li/hi-go-server/model"
	"github.com/yanlong-li/hi-go-server/packet_model/server/channel"
	"github.com/yanlong-li/hi-go-server/packet_model/trait"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
)

func init() {
	route.Register(channel.Get{}, func(info channel.Get, conn connect.Connector) {
		_cha := &model.Channel{}
		ormErr := db.Model(_cha).Find().Where("=", "id", info.Id).One()
		if !ormErr.Status() {
			_ = conn.Send(channel.GetFail{Fail: trait.Fail{Message: "数据不存在"}})
		} else {
			createUser, _ := model.GetUserById(_cha.CreateUserId)
			ownerUser, _ := model.GetUserById(_cha.OwnerUserId)
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
				Avatar:     _cha.Avatar,
				Describe:   _cha.Describe,
				Channels:   []channel.Info{},
			}
			ChannelChildrenS := _cha.GetChildren()
			for _, channelChildren := range ChannelChildrenS {
				if _channelChildren, ok := channelChildren.(model.Channel); ok {
					_createUser, _ := model.GetUserById(_cha.CreateUserId)
					_ownerUser, _ := model.GetUserById(_cha.OwnerUserId)
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

			_ = conn.Send(info)
		}
	})
}
