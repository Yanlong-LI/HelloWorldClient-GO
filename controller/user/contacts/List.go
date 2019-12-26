package contacts

import (
	"HelloWorld/io/db"
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/route"
	"HelloWorldServer/model"
	conn2 "HelloWorldServer/model/online"
	"HelloWorldServer/packet/user/contacts"
)

func init() {
	route.Register(contacts.GetList{}, actionGetList)
}

func actionGetList(list contacts.GetList, conn connect.Connector) {
	_list := contacts.List{}
	_list.List = make([]contacts.Info, 0)
	selfUser, _ := conn2.Auth(conn.GetId())

	userContacts := db.Model(&model.UserContact{}).Where("=", "user_id", selfUser.Id).All()

	for _, contact := range userContacts {
		if _contact, ok := contact.(model.UserContact); ok {
			_contactInfo, err := _contact.GetContactInfo()
			if err == nil {
				_contact := contacts.Info{
					Id:       _contactInfo.Id,
					Nickname: _contactInfo.Nickname,
					Avatar:   _contactInfo.Avatar,
					Language: _contactInfo.Language,
					Region:   _contactInfo.Region,
					Remarks:  _contact.Remarks,
				}
				_list.List = append(_list.List, _contact)
			}
		}

	}
	conn.Send(_list)

}
