package channel

import (
	"github.com/yanlong-li/hi-go-server/common"
	"github.com/yanlong-li/hi-go-server/model"
	"github.com/yanlong-li/hi-go-server/packet_model/server/channel"
	"github.com/yanlong-li/hi-go-server/packet_model/trait"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
	"strings"
	"time"
)

func init() {
	route.Register(channel.AddChannel{}, actionAddChannel)
	route.Register(channel.AddSubChannel{}, actionAddSubChannel)
}

func actionAddChannel(addChannel channel.AddChannel, conn connect.Connector) {

	if len(strings.Trim(addChannel.Name, " 	")) == 0 {
		_ = conn.Send(channel.AddChannelFail{Fail: trait.Fail{Message: "名称不能为空"}})
		return
	}

	_ = conn.Send(channel.AddChannelFail{Fail: trait.Fail{Message: "目前不开放主频道创建"}})
	return

	// 创建频道
	timeNow := uint64(time.Now().Unix())
	userId, _ := common.Auth(conn.GetId())
	newChannel := &model.Channel{Name: addChannel.Name, Avatar: addChannel.Avatar, CreateTime: timeNow,
		UpdateTime: timeNow, CreateUserId: userId.Id, OwnerUserId: userId.Id, Status: 1, ParentId: 0, ServerId: 1,
		Describe: addChannel.Describe,
	}
	err := db.Model(newChannel).Insert().Insert()
	if !err.Status() {
		_ = conn.Send(channel.AddChannelFail{Fail: trait.Fail{Message: err.Error()}})
		return
	}
	// 将当前用户加入频道
	newChannelUser := &model.ChannelUser{ChannelId: newChannel.Id, UserId: userId.Id, CreateTime: timeNow, UpdateTime: timeNow, OpenId: ""}
	err = db.Model(newChannelUser).Insert().Insert()
	if !err.Status() {
		_ = conn.Send(channel.AddChannelFail{Fail: trait.Fail{Message: err.Error()}})
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
	_ = conn.Send(channel.AddChannelSuccess{Info: channelInfo})
}

// 添加子频道
func actionAddSubChannel(addChannel channel.AddSubChannel, conn connect.Connector) {
	userInfo, _ := common.Auth(conn.GetId())
	if len(strings.Trim(addChannel.Name, " 	")) == 0 {
		_ = conn.Send(channel.AddSubChannelFail{Fail: trait.Fail{Message: "名称不能为空"}})
		return
	}
	// 查询频道是否存在

	_channel := &model.Channel{}
	_exists := db.Model(_channel).Find().Where("id", addChannel.ChannelId).Exists()
	if !_exists {
		_ = conn.Send(channel.AddSubChannelFail{Fail: trait.Fail{Message: "频道不存在"}})
		return
	}

	if _channel.OwnerUserId != userInfo.Id {
		_ = conn.Send(channel.AddSubChannelFail{Fail: trait.Fail{Message: "您不是OW，没有权限"}})
		return
	}

	// 创建频道
	timeNow := uint64(time.Now().Unix())

	newChannel := &model.Channel{Name: addChannel.Name, Avatar: addChannel.Avatar, CreateTime: timeNow,
		UpdateTime: timeNow, CreateUserId: userInfo.Id, OwnerUserId: userInfo.Id, Status: 1, ParentId: addChannel.ChannelId, ServerId: 1,
		Describe: addChannel.Describe,
	}
	err := db.Model(newChannel).Insert().Insert()
	if !err.Status() {
		_ = conn.Send(channel.AddSubChannelFail{Fail: trait.Fail{Message: err.Error()}})
		return
	}

	channelInfo := channel.Info{Id: newChannel.Id, Name: newChannel.Name, Avatar: newChannel.Avatar, CreateUser: struct {
		Id       uint64
		Nickname string
	}{Id: userInfo.Id, Nickname: userInfo.Nickname}, Describe: newChannel.Describe, OwnerUser: struct {
		Id       uint64
		Nickname string
	}{Id: userInfo.Id, Nickname: userInfo.Nickname}, CreateTime: timeNow, Public: true, Verify: true, Commerce: true, Channels: make([]channel.Info, 0)}
	// 返回当前频道信息
	_ = conn.Send(channel.AddSubChannelSuccess{Channel: channelInfo, ChannelId: addChannel.ChannelId})
}
