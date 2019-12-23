package contacts

import (
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/route"
	conn2 "HelloWorldServer/model/Login"
	"HelloWorldServer/packet/user/contacts"
)

func init() {
	route.Register(contacts.GetList{}, actionGetList)
}

func actionGetList(list contacts.GetList, conn connect.Connector) {
	_list := contacts.List{}
	_list.List = make([]contacts.Info, 0)
	selfUser, _ := conn2.Auth(conn.GetId())
	for _, _userOnline := range conn2.Login {

		if _userOnline.Id == selfUser.Id {
			continue
		}
		_contact := contacts.Info{Id: _userOnline.Id, Nickname: _userOnline.Nickname, Avatar: _userOnline.Avatar, Language: _userOnline.Language, Region: _userOnline.Region}
		_list.List = append(_list.List, _contact)

	}
	conn.Send(_list)

}
