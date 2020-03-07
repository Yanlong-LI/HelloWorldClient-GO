package channel

import (
	"github.com/yanlong-li/HelloWorld-GO/io/db"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorldServer/model"
	"github.com/yanlong-li/HelloWorldServer/model/online"
	"github.com/yanlong-li/HelloWorldServer/packetModel/server/channel"
	"github.com/yanlong-li/HelloWorldServer/packetModel/trait"
	"strings"
	"time"
)

func init() {
	route.Register(channel.AddChannel{}, actionAddChannel)
	route.Register(channel.AddSubChannel{}, actionAddSubChannel)
}

func actionAddChannel(addChannel channel.AddChannel, conn connect.Connector) {

	if len(strings.Trim(addChannel.Name, " 	")) == 0 {
		conn.Send(channel.AddChannelFail{Fail: trait.Fail{Message: "名称不能为空"}})
		return
	}

	// 创建频道
	timeNow := uint64(time.Now().Unix())
	userId, _ := online.Auth(conn.GetId())
	newChannel := &model.Channel{Name: addChannel.Name, Avatar: addChannel.Avatar, CreateTime: timeNow,
		UpdateTime: timeNow, CreateUserId: userId.Id, OwnerUserId: userId.Id, Status: 1, ParentId: 0, ServerId: 1,
		Describe: addChannel.Describe,
	}
	err := db.Model(newChannel).Insert()
	if err != nil {
		conn.Send(channel.AddChannelFail{Fail: trait.Fail{Message: err.Error()}})
		return
	}
	// 将当前用户加入频道
	newChannelUser := &model.ChannelUser{ChannelId: newChannel.Id, UserId: userId.Id, CreateTime: timeNow, UpdateTime: timeNow, OpenId: ""}
	err = db.Model(newChannelUser).Insert()
	if err != nil {
		conn.Send(channel.AddChannelFail{Fail: trait.Fail{Message: err.Error()}})
		return
	}

	channelInfo := channel.Info{Id: newChannel.Id, Name: newChannel.Name, Avatar: newChannel.Avatar, CreateUser: struct {
		Id       uint64
		Nickname string
	}{Id: userId.Id, Nickname: userId.Nickname}, Describe: newChannel.Describe, OwnerUser: struct {
		Id       uint64
		Nickname string
	}{Id: userId.Id, Nickname: userId.Nickname}, CreateTime: timeNow, Public: true, Verify: true, Commerce: true, Channels: make([]channel.Info, 0)}
	// 返回当前频道信息
	conn.Send(channel.AddChannelSuccess{Info: channelInfo})
}

// 添加子频道
func actionAddSubChannel(addChannel channel.AddSubChannel, conn connect.Connector) {

	if len(strings.Trim(addChannel.Name, " 	")) == 0 {
		conn.Send(channel.AddSubChannelFail{Fail: trait.Fail{Message: "名称不能为空"}})
		return
	}
	// 查询频道是否存在

	_channel := &model.Channel{}
	_exists := db.Model(_channel).Where("id", addChannel.ChannelId).Exists()
	if !_exists {
		conn.Send(channel.AddSubChannelFail{Fail: trait.Fail{Message: "频道不存在"}})
		return
	}

	// 创建频道
	timeNow := uint64(time.Now().Unix())
	userId, _ := online.Auth(conn.GetId())
	newChannel := &model.Channel{Name: addChannel.Name, Avatar: addChannel.Avatar, CreateTime: timeNow,
		UpdateTime: timeNow, CreateUserId: userId.Id, OwnerUserId: userId.Id, Status: 1, ParentId: addChannel.ChannelId, ServerId: 1,
		Describe: addChannel.Describe,
	}
	err := db.Model(newChannel).Insert()
	if err != nil {
		conn.Send(channel.AddSubChannelFail{Fail: trait.Fail{Message: err.Error()}})
		return
	}

	channelInfo := channel.Info{Id: newChannel.Id, Name: newChannel.Name, Avatar: newChannel.Avatar, CreateUser: struct {
		Id       uint64
		Nickname string
	}{Id: userId.Id, Nickname: userId.Nickname}, Describe: newChannel.Describe, OwnerUser: struct {
		Id       uint64
		Nickname string
	}{Id: userId.Id, Nickname: userId.Nickname}, CreateTime: timeNow, Public: true, Verify: true, Commerce: true, Channels: make([]channel.Info, 0)}
	// 返回当前频道信息
	conn.Send(channel.AddSubChannelSuccess{Channel: channelInfo, ChannelId: addChannel.ChannelId})
}