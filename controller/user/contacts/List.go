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
	_user, _ := conn2.Auth(conn.GetId())
	for _, __user := range conn2.Login {

		if __user.Id == _user.Id {
			continue
		}
		_contact := contacts.Info{Id: __user.Id, UserName: __user.UserName, Avatar: __user.Avatar, Language: __user.Language, Region: __user.Region}
		_list.List = append(_list.List, _contact)

	}
	conn.Send(_list)

}
