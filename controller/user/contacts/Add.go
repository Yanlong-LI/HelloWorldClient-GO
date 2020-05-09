package contacts

import (
	"github.com/yanlong-li/HelloWorld-GO/io/db"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorldServer/model"
	conn2 "github.com/yanlong-li/HelloWorldServer/model/online"
	"github.com/yanlong-li/HelloWorldServer/packetModel/trait"
	"github.com/yanlong-li/HelloWorldServer/packetModel/user"
	"github.com/yanlong-li/HelloWorldServer/packetModel/user/contacts"
	"time"
)

func init() {
	route.Register(searchUser, searchUser)
}

func searchUser(searchUser contacts.SearchUser, connector connect.Connector) {
	_userModel := &model.User{}
	err := db.Model(_userModel).Find().Where("=", "account", searchUser.Account).One()
	if !err.Status() || err.Empty() {
		_ = connector.Send(contacts.SearchUserFail{Fail: struct {
			Code    uint32
			Message string
		}{Code: uint32(1), Message: "未搜索到账户"}})
		return
	}
	_userInfo := user.Info{Id: _userModel.Id, Nickname: _userModel.Nickname, Avatar: _userModel.Avatar, Region: _userModel.Region, Language: _userModel.Language}
	_ = connector.Send(contacts.SearchUserSuccess{
		Info: _userInfo,
	})
}

func addContact(AddContact contacts.AddContact, connector connect.Connector) {

	selfUser, _ := conn2.Auth(connector.GetId())

	////查询当前是否是好友
	_contactUser := &model.UserContact{}
	_contactUser2 := &model.UserContact{}
	//查询当前是否被对方拉黑
	ormErr := db.Model(_contactUser).Find().Where(map[interface{}]interface{}{"user_id": AddContact.Id, "contact_id": selfUser.Id}).One()
	ormErr2 := db.Model(_contactUser2).Find().Where(map[interface{}]interface{}{"user_id": AddContact.Id, "contact_id": selfUser.Id}).One()
	if !ormErr.Empty() {
		//判断是否在黑名单中
		if _contactUser.Status&2 == 2 {
			//在黑名单
			_ = connector.Send(contacts.AddContactFail{
				Fail: trait.Fail{Message: "申请发送失败"},
			})
			return
		}
		if _contactUser.Status&1 == 1 {
			//已经是好友了，直接添加
			//_ = connector.Send(contacts.AddContactFail{
			//	Fail: trait.Fail{Message: "申请发送失败"},
			//})
			if (ormErr2).Empty() {
				_contactUser2.UserId = selfUser.Id
				_contactUser2.ContactId = AddContact.Id
				_contactUser2.Status = 1
				_contactUser2.CreateTime = uint64(time.Now().Unix())
				_contactUser2.UpdateTime = _contactUser2.CreateTime
				_ = db.Model(_contactUser2).Insert().Insert()
				_ = connector.Send(contacts.AcceptedContact{Id: AddContact.Id})
			} else {
				// 存在？ 黑名单、删除、待接受什么的都设置为正常 1
				_contactUser2.Status = 1
				_contactUser2.UpdateTime = uint64(time.Now().Unix())
				_, _ = db.Model(_contactUser2).Update().Update()
				_ = connector.Send(contacts.AcceptedContact{Id: AddContact.Id})
			}

			return
		}
		//其它的要重新申请
	}

	if ormErr2.Empty() {
		//添加进好友列表，并通知对方申请消息
		_contactUser2.UserId = selfUser.Id
		_contactUser2.ContactId = AddContact.Id
		_contactUser2.Status = 8
		_contactUser2.CreateTime = uint64(time.Now().Unix())
		_contactUser2.UpdateTime = _contactUser2.CreateTime
		_ = db.Model(_contactUser2).Insert().Insert()
	} else {
		// 存在？ 黑名单、删除、待接受什么的都设置为正常 1
		_contactUser2.Status = 8
		_contactUser2.UpdateTime = uint64(time.Now().Unix())
		_, _ = db.Model(_contactUser2).Update().Update()
	}

	if ormErr.Empty() {
		//添加进好友列表，并通知对方申请消息
		_contactUser2.UserId = AddContact.Id
		_contactUser2.ContactId = selfUser.Id
		_contactUser2.Status = 0
		_contactUser2.CreateTime = uint64(time.Now().Unix())
		_contactUser2.UpdateTime = _contactUser2.CreateTime
		_ = db.Model(_contactUser2).Insert().Insert()
	} else {
		// 存在？ 黑名单、删除、待接受什么的都设置为正常 1
		_contactUser.Status = 0
		_contactUser.UpdateTime = uint64(time.Now().Unix())
		_, _ = db.Model(_contactUser2).Update().Update()
	}

	_ = connector.Send(contacts.AddContactSuccess{})

	// 转换身份
	AddContact.Id = selfUser.Id
	conn2.UserSendMessage(AddContact.Id, contacts.RequestAddContact{AddContact: AddContact})
}

func acceptContact(contact contacts.AcceptContact, connector connect.Connector) {
	//todo 查询当前是否是好友
	//todo 查询当前状态
	//todo 添加对方为好友，并通知对方已通过
}
func refuseContact(contact contacts.RefuseContact, connector connect.Connector) {
	//todo 查询当前是否是好友
	//todo 查询当前状态
	//todo 修改对方的申请状态，通知对方被拒绝
}
