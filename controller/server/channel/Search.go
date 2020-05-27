package channel

import (
	"github.com/yanlong-li/HelloWorld-GO/io/db"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorldServer/model"
	"github.com/yanlong-li/HelloWorldServer/packetModel/server/channel"
)

func init() {
	route.Register(channel.SearchChannelList{}, actionSearchChannelList)
}

func actionSearchChannelList(searchList channel.SearchChannelList, conn connect.Connector) {

	_list := db.Model(&model.Channel{}).Find().Where("like", "name", "%"+searchList.Name+"%").AndWhere("=", "parent_id", 0).AndWhere("=", "delete_time", 0).All()

	list := channel.SearchChannelListSuccess{}
	list.List = make([]channel.Info, 0)
	for _, cha := range _list {
		if _cha, ok := cha.(model.Channel); ok {

			createUser, _ := model.GetUserById(_cha.CreateUserId)
			ownerUser, _ := model.GetUserById(_cha.CreateUserId)

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
			}
			list.List = append(list.List, info)
		}
	}

	_ = conn.Send(list)

}
