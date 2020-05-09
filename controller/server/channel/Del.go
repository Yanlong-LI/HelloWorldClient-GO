package channel

import (
	"github.com/yanlong-li/HelloWorld-GO/io/db"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorldServer/model"
	"github.com/yanlong-li/HelloWorldServer/model/online"
	"github.com/yanlong-li/HelloWorldServer/packetModel/server/channel"
	"github.com/yanlong-li/HelloWorldServer/packetModel/trait"
)

func init() {
	route.Register(channel.DelChannel{}, actionDelChannel)
}

func actionDelChannel(delChannel channel.DelChannel, conn connect.Connector) {

	var _channel = &model.Channel{}
	err := db.Model(_channel).Find().Where("id", delChannel.Id).One()
	if err.Empty() {
		_ = conn.Send(channel.DelChannelFail{Fail: trait.Fail{Message: err.Error()}})
		return
	}
	user, _ := online.Auth(conn.GetId())
	if _channel.OwnerUserId != user.Id {
		_ = conn.Send(channel.DelChannelFail{Fail: trait.Fail{Message: "您未拥有权限"}})
		return
	}
	//删除频道的用户
	_, err = db.Model(&model.ChannelUser{}).Delete().Where("channel_id", _channel.Id).Delete()
	// 删除子频道
	_, err = db.Model(_channel).Delete().Where("parent_id", _channel.Id).Delete()
	// 删除频道
	_, err = db.Model(_channel).Delete().Where("id", _channel.Id).Delete()
	if err.Empty() {
		_ = conn.Send(channel.DelChannelFail{Fail: trait.Fail{Message: err.Error()}})
		return
	}
	_ = conn.Send(channel.DelChannelSuccess{Id: _channel.Id})

}
