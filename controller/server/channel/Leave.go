package channel

import (
	"github.com/yanlong-li/HelloWorld-GO/io/db"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorldServer/model"
	"github.com/yanlong-li/HelloWorldServer/model/online"
	"github.com/yanlong-li/HelloWorldServer/packetModel/server/channel"
	"github.com/yanlong-li/HelloWorldServer/packetModel/trait"
	"time"
)

func init() {
	route.Register(channel.LeaveChannel{}, actionLeaveChannel)
}

func actionLeaveChannel(leaveChannel channel.LeaveChannel, conn connect.Connector) {
	user, _ := online.Auth(conn.GetId())

	var _channelUser = &model.ChannelUser{}
	err := db.Model(_channelUser).Find().Where(map[interface{}]interface{}{
		"channel_id": leaveChannel.Id,
		"user_id":    user.Id,
	}).One()
	if err.Empty() {
		_ = conn.Send(channel.LeaveChannelFail{Fail: trait.Fail{Message: "未查到数据"}})
		return
	}

	_channelUser.DeleteTime = uint64(time.Now().Unix())
	_channelUser.UpdateTime = _channelUser.DeleteTime
	_, err = db.Model(_channelUser).Update().Update()

	if !err.Status() {
		_ = conn.Send(channel.DelChannelFail{Fail: trait.Fail{Message: err.Error()}})
		return
	}
	_ = conn.Send(channel.DelChannelSuccess{Id: _channelUser.Id})

}
