package common

import (
	logger "github.com/yanlong-li/hi-go-logger"
	db "github.com/yanlong-li/hi-go-orm"
	"github.com/yanlong-li/hi-go-server/model"
)

func BroadcastToChannel(channelId, userId uint64, message interface{}) {
	//查询频道用户列表
	userList := db.Model(&model.ChannelUser{}).Find().Where(map[interface{}]interface{}{"delete_time": 0, "channel_id": channelId}).All()
	//针对用户发送数据
	for _, v := range userList {
		_cu, ok := v.(model.ChannelUser)
		if ok != true {
			logger.Fatal("断言错误", 0, v)
			continue
		}
		// 不发送给自己
		if _cu.UserId == userId {
			continue
		}
		SendMessageToUser(_cu.UserId, message)
	}
}
