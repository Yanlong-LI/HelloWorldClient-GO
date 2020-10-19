package channel

import (
	db "github.com/yanlong-li/hi-go-orm"
	"github.com/yanlong-li/hi-go-server/common"
	"github.com/yanlong-li/hi-go-server/model"
	"github.com/yanlong-li/hi-go-server/packet_model/server/channel"
	"github.com/yanlong-li/hi-go-server/packet_model/trait"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
	"time"
)

func init() {
	route.Register(channel.LeaveChannel{}, actionLeaveChannel)
}

func actionLeaveChannel(leaveChannel channel.LeaveChannel, conn connect.Connector) {
	user, _ := common.Auth(conn.GetId())

	var _channelUser = &model.ChannelUser{}
	err := db.Model(_channelUser).Find().Where(map[interface{}]interface{}{
		"channel_id":  leaveChannel.Id,
		"user_id":     user.Id,
		"delete_time": 0,
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
