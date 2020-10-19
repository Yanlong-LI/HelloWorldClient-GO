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
	route.Register(channel.DelChannel{}, actionDelChannel)
}

func actionDelChannel(delChannel channel.DelChannel, conn connect.Connector) {

	var _channel = &model.Channel{}
	err := db.Model(_channel).Find().Where("id", delChannel.Id).One()
	if err.Empty() {
		_ = conn.Send(channel.DelChannelFail{Fail: trait.Fail{Message: err.Error()}})
		return
	}
	user, _ := common.Auth(conn.GetId())
	if _channel.OwnerUserId != user.Id {
		_ = conn.Send(channel.DelChannelFail{Fail: trait.Fail{Message: "您未拥有权限"}})
		return
	}
	////删除频道的用户
	//_, err = db.Model(&model.ChannelUser{}).Delete().Where("channel_id", _channel.Id).Delete()
	//// 删除子频道
	//_, err = db.Model(_channel).Delete().Where("parent_id", _channel.Id).Delete()
	//// 删除频道
	//_, err = db.Model(_channel).Delete().Where("id", _channel.Id).Delete()
	_channel.DeleteTime = uint64(time.Now().Unix())
	_channel.UpdateTime = _channel.DeleteTime
	_, err = db.Model(_channel).Update().Update()

	if !err.Status() {
		_ = conn.Send(channel.DelChannelFail{Fail: trait.Fail{Message: err.Error()}})
		return
	}
	_ = conn.Send(channel.DelChannelSuccess{Id: _channel.Id})

}
