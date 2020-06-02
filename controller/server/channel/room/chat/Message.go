package chat

import (
	"github.com/yanlong-li/HelloWorld-GO/io/db"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorldServer/common"
	"github.com/yanlong-li/HelloWorldServer/model"
	"github.com/yanlong-li/HelloWorldServer/packetModel/server/channel/room/message"
	"github.com/yanlong-li/HelloWorldServer/packetModel/trait"
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

	var _channelUser = &model.ChannelUser{}
	ormErr := db.Model(_channelUser).Find().Where(map[interface{}]interface{}{
		"channel_id":  msg.ChannelId,
		"user_id":     _user.Id,
		"delete_time": 0,
	}).One()
	if ormErr.Empty() {
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
