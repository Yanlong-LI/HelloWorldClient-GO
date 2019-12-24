package channel

import (
	"HelloWorld/io/db"
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/route"
	"HelloWorldServer/model"
	"HelloWorldServer/packet/server/channel"
)

func init() {
	route.Register(channel.SearchChannelList{}, actionSearchChannelList)
}

func actionSearchChannelList(searchList channel.SearchChannelList, conn connect.Connector) {

	_list := db.Find(&model.Channel{}).Where("like", "name", "%"+searchList.Name+"%").All()

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
				Icon:       _cha.Icon,
			}
			list.List = append(list.List, info)
		}
	}

	conn.Send(list)

}
