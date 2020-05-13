package contacts

import (
	"github.com/yanlong-li/HelloWorld-GO/io/db"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorldServer/model"
	conn2 "github.com/yanlong-li/HelloWorldServer/model/online"
	"github.com/yanlong-li/HelloWorldServer/packetModel/user/contacts"
)

func init() {
	route.Register(contacts.GetList{}, actionGetList)
	route.Register(contacts.GetRequestList{}, actionGetRequestList)
	route.Register(contacts.GetBlacklist{}, actionGetBlackList)
}

func actionGetList(list contacts.GetList, conn connect.Connector) {
	_list := contacts.List{}
	_list.List = make([]contacts.Info, 0)
	selfUser, _ := conn2.Auth(conn.GetId())

	userContacts := db.Model(&model.UserContact{}).Find().Where("=", "user_id", selfUser.Id).All()

	for _, contact := range userContacts {
		if _contact, ok := contact.(model.UserContact); ok {
			_contactInfo, err := _contact.GetContactUser()
			if err.Status() {
				_contact := contacts.Info{
					Id:       _contactInfo.Id,
					Nickname: _contactInfo.Nickname,
					Avatar:   _contactInfo.Avatar,
					Language: _contactInfo.Language,
					Region:   _contactInfo.Region,
					Remark:   _contact.Remark,
				}
				_list.List = append(_list.List, _contact)
			}
		}

	}
	_ = conn.Send(_list)

}

func actionGetRequestList(list contacts.GetList, conn connect.Connector) {
	_list := contacts.RequestList{}

	selfUser, _ := conn2.Auth(conn.GetId())

	userContacts := db.Model(&model.UserContactRequest{}).Find().Where("=", "contact_id", selfUser.Id).All()

	for _, contact := range userContacts {
		if _contact, ok := contact.(model.UserContactRequest); ok {
			_contactInfo, err := _contact.GetUser()
			if err.Status() {
				_contact := contacts.Info{
					Id:       _contactInfo.Id,
					Nickname: _contactInfo.Nickname,
					Avatar:   _contactInfo.Avatar,
					Language: _contactInfo.Language,
					Region:   _contactInfo.Region,
					Remark:   _contact.UserRemark,
				}
				_list.List = append(_list.List, _contact)
			}
		}

	}
	_ = conn.Send(_list)

}

func actionGetBlackList(list contacts.GetList, conn connect.Connector) {
	_list := contacts.Blacklist{}
	selfUser, _ := conn2.Auth(conn.GetId())

	userContactBlacks := db.Model(&model.UserContactBlack{}).Find().Where("=", "user_id", selfUser.Id).All()

	for _, userContactBlack := range userContactBlacks {
		if contactBlack, ok := userContactBlack.(model.UserContactBlack); ok {
			user, err := contactBlack.GetContactUser()
			if err.Status() {
				_contact := contacts.Info{
					Id:       user.Id,
					Nickname: user.Nickname,
					Avatar:   user.Avatar,
					Language: user.Language,
					Region:   user.Region,
					Remark:   contactBlack.Remark,
				}
				_list.List = append(_list.List, _contact)
			}
		}

	}
	_ = conn.Send(_list)

}
