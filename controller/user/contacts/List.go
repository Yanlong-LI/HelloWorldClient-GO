package contacts

import (
	db "github.com/yanlong-li/hi-go-orm"
	"github.com/yanlong-li/hi-go-server/common"
	"github.com/yanlong-li/hi-go-server/model"
	"github.com/yanlong-li/hi-go-server/packet_model/user/contacts"
	"github.com/yanlong-li/hi-go-socket/connect"
	"github.com/yanlong-li/hi-go-socket/route"
)

func init() {
	route.Register(contacts.GetList{}, actionGetList)
	route.Register(contacts.GetRequestList{}, actionGetRequestList)
	route.Register(contacts.GetBlacklist{}, actionGetBlackList)
}

// 获取联系人列表
func actionGetList(list contacts.GetList, conn connect.Connector) {
	var _list contacts.List = make([]contacts.Info, 0)
	selfUser, _ := common.Auth(conn.GetId())

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
					Online:   common.UserOnlineByUserId(_contactInfo.Id),
				}
				_list = append(_list, _contact)
			}
		}

	}
	_ = conn.Send(_list)

}

// 获取请求列表
func actionGetRequestList(list contacts.GetList, conn connect.Connector) {
	_list := contacts.RequestList{}

	selfUser, _ := common.Auth(conn.GetId())

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
					Online:   common.UserOnlineByUserId(_contactInfo.Id),
				}
				_list = append(_list, _contact)
			}
		}

	}
	_ = conn.Send(_list)

}

func actionGetBlackList(list contacts.GetList, conn connect.Connector) {
	_list := contacts.Blacklist{}
	selfUser, _ := common.Auth(conn.GetId())

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
					Online:   false, // 给名单的用户保持离线状态返回
				}
				_list = append(_list, _contact)
			}
		}

	}
	_ = conn.Send(_list)

}
