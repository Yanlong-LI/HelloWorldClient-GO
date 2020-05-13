package contacts

import (
	"github.com/yanlong-li/HelloWorld-GO/io/db"
	"github.com/yanlong-li/HelloWorld-GO/io/logger"
	"github.com/yanlong-li/HelloWorld-GO/io/network/connect"
	"github.com/yanlong-li/HelloWorld-GO/io/network/route"
	"github.com/yanlong-li/HelloWorldServer/model"
	conn2 "github.com/yanlong-li/HelloWorldServer/model/online"
	"github.com/yanlong-li/HelloWorldServer/packetModel/trait"
	"github.com/yanlong-li/HelloWorldServer/packetModel/user/contacts"
	"time"
)

func init() {
	route.Register(contacts.AddContact{}, addContact)
	route.Register(contacts.AcceptContact{}, acceptContact)
	route.Register(contacts.RefuseContact{}, refuseContact)
}

func addContact(_addContact contacts.AddContact, connector connect.Connector) {

	var ormErr db.OrmError
	_selfUser, _ := conn2.Auth(connector.GetId())

	if _selfUser.Id == _addContact.Id {
		_ = connector.Send(contacts.AddContactFail{
			Fail: trait.Fail{Message: "不能添加自己"},
		})
	}

	// 查询是否被对方拉黑
	blackUser := &model.UserContactBlack{}
	ormErr = db.Model(blackUser).Find().Where(map[interface{}]interface{}{"user_id": _addContact.Id, "contact_id": _selfUser.Id}).One()
	if ormErr.Status() && !ormErr.Empty() {
		//在黑名单
		_ = connector.Send(contacts.AddContactFail{
			Fail: trait.Fail{Message: "您进入了对方的黑名单"},
		})
	}
	_contactUser := &model.UserContact{}

	ormErr = db.Model(_contactUser).Find().Where(map[interface{}]interface{}{"user_id": _addContact.Id, "contact_id": _selfUser.Id}).One()
	if !ormErr.Empty() {
		model.UserContactAddUser(_selfUser.Id, _addContact.Id)
		//发送给自己请求通过
		_ = connector.Send(contacts.AcceptedContact{Id: _addContact.Id})
		return
	}
	// 如果为空，需要申请
	_userContactRequest := &model.UserContactRequest{}
	ormErr = db.Model(_userContactRequest).Find().Where(map[interface{}]interface{}{"user_id": _selfUser.Id, "contact_id": _addContact.Id}).One()
	if ormErr.Empty() {
		// 创建新请求
		_userContactRequest.UserId = _selfUser.Id
		_userContactRequest.ContactId = _addContact.Id
		_userContactRequest.UserRemark = _addContact.Remark
		_userContactRequest.CreateTime = uint64(time.Now().Unix())
		_userContactRequest.UpdateTime = uint64(time.Now().Unix())
		_userContactRequest.Status = model.UserContactRequestStatusRequest
		// 插入数据库
		ormErr = db.Model(_userContactRequest).Insert().Insert()
		if !ormErr.Status() {
			logger.Warning("创建 联系人请求 写入数据库失败:"+ormErr.Error(), 0, ormErr)
		}
	} else {

		if _userContactRequest.UpdateTime+600 >= uint64(time.Now().Unix()) {
			_ = connector.Send(contacts.AddContactFail{
				Fail: trait.Fail{Message: "操作频繁，请稍后再试"},
			})
			return
		}

		// 更新请求
		_userContactRequest.UpdateTime = uint64(time.Now().Unix())
		_userContactRequest.Status = model.UserContactRequestStatusRequest
		_, ormErr = db.Model(_userContactRequest).Update().Update()
		if !ormErr.Status() {
			logger.Warning("更新 联系人请求 写入数据库失败:"+ormErr.Error(), 0, ormErr)
		}
	}

	_ = connector.Send(contacts.AddContactSuccess{})
	// 通知目标有新的请求
	conn2.UserSendMessage(_addContact.Id, contacts.RequestAddContact{
		AddContact: contacts.AddContact{Id: _selfUser.Id, Remark: _addContact.Remark},
	})
}

func acceptContact(_acceptContact contacts.AcceptContact, conn connect.Connector) {

	selfUser, _ := conn2.Auth(conn.GetId())

	ucr, ormErr := model.UserContactRequestFind(_acceptContact.Id, selfUser.Id)
	if ormErr.Empty() {
		_ = conn.Send(contacts.AcceptContactFail{Fail: trait.Fail{Message: "请求不存在"}})
		return
	}
	if ucr.Status != model.UserContactRequestStatusRequest {
		_ = conn.Send(contacts.AcceptContactFail{Fail: trait.Fail{Message: "请求已被处理"}})
		return
	}
	// 标记为已处理
	ucr.Status = model.UserContactRequestStatusSuccess
	_ = ucr.Update()

	model.UserContactAddUser(selfUser.Id, _acceptContact.Id)
	model.UserContactAddUser(_acceptContact.Id, selfUser.Id)

	_ = conn.Send(contacts.AcceptContactSuccess{Success: trait.Success{}})

	// 通知对方我已通过
	conn2.UserSendMessage(_acceptContact.Id, contacts.AcceptedContact{Id: selfUser.Id})

}
func refuseContact(_contact contacts.RefuseContact, conn connect.Connector) {
	selfUser, _ := conn2.Auth(conn.GetId())

	ucr, ormErr := model.UserContactRequestFind(_contact.Id, selfUser.Id)
	if ormErr.Empty() {
		_ = conn.Send(contacts.RefuseContactFail{Fail: trait.Fail{Message: "请求不存在"}})
		return
	}
	if ucr.Status != model.UserContactRequestStatusRequest {
		_ = conn.Send(contacts.RefuseContactFail{Fail: trait.Fail{Message: "请求已被处理"}})
		return
	}
	// 标记为已处理
	ucr.Status = model.UserContactRequestStatusRefuse
	_ = ucr.Update()

	_ = conn.Send(contacts.RefuseContactSuccess{Success: trait.Success{}})

	// 通知对方我已通过
	conn2.UserSendMessage(_contact.Id, contacts.RejectedContact{Id: selfUser.Id, Remark: _contact.Remark})
}
