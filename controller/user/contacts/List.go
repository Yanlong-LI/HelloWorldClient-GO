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

	for _, _user := range conn2.Login {

		_contact := contacts.Info{Id: _user.Id, UserName: _user.UserName, Avatar: _user.Avatar, Language: _user.Language, Region: _user.Region}
		_list.List = append(_list.List, _contact)

	}
	conn.Send(_list)

}
