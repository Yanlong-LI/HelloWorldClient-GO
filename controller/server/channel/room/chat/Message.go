package chat

import (
	"github.com/yanlong-li/hi-go-server/common"
	"github.com/yanlong-li/hi-go-server/model"
	"github.com/yanlong-li/hi-go-server/packet_model/server/channel/room/message"
	"github.com/yanlong-li/hi-go-server/packet_model/trait"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
	"log"
	"time"
)

func init() {
	route.Register(message.SendTextMessage{}, TextMessage)
}

func TextMessage(msg message.SendTextMessage, conn connect.Connector) {

	_user, err := common.Auth(conn.GetId())
	if err != nil {

		log.Print("收到用户消息：获取用户错误")
		return
	}

	if !checkUserJoinChannel(_user.Id, msg.ChannelId) {
		_ = conn.Send(message.SendTextMessageFail{Fail: trait.Fail{Message: "您没有权限发送数据"}, ChannelId: msg.ChannelId, ServerId: msg.ServerId, RandomString: msg.RandomString})
		return
	}

	_msg := message.TextMessage{
		SendTextMessage: msg,
		Time:            uint64(time.Now().Unix()),
		Author: struct {
			Id       uint64
			Nickname string
		}{Id: _user.Id, Nickname: _user.Nickname},
	}

	_ = conn.Send(message.SendTextMessageSuccess{TextMessage: _msg})

	common.BroadcastToChannel(msg.ChannelId, _user.Id, _msg)
}

func checkUserJoinChannel(userId, channelId uint64) bool {
	//先查频道是否存在
	_channel := &model.Channel{}
	ormErr := db.Model(_channel).Find().Where(map[interface{}]interface{}{
		"id":          channelId,
		"delete_time": 0,
	}).One()

	if !ormErr.Status() {
		return false
	}
	// 再查是否在当前目录
	_channelUser := &model.ChannelUser{}
	ormErr = db.Model(_channelUser).Find().Where(map[interface{}]interface{}{
		"channel_id":  channelId,
		"user_id":     userId,
		"delete_time": 0,
	}).One()

	if ormErr.Status() {
		return true
	}
	// 查询父频道
	if _channel.ParentId != 0 {
		return checkUserJoinChannel(userId, _channel.ParentId)
	}

	return false
}
