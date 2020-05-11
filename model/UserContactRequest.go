package model

import (
	"github.com/yanlong-li/HelloWorld-GO/io/db"
	"time"
)

type UserContactRequest struct {
	Id            uint64
	UserId        uint64
	UserRemark    string
	ContactId     uint64
	ContactRemark string
	Status        uint16
	CreateTime    uint64
	UpdateTime    uint64
}

func (uc *UserContactRequest) GetUserInfo() (User, db.OrmError) {
	return GetUserById(uc.UserId)
}

func (uc *UserContactRequest) GetContactInfo() (User, db.OrmError) {
	return GetUserById(uc.ContactId)
}

func UserContactRequestFind(userId, contactId uint64) (*UserContactRequest, db.OrmError) {

	userContactRequest := &UserContactRequest{}
	ormErr := db.Model(userContactRequest).Find().Where("=", "user_id", userId).AndWhere("=", "contact_id", contactId).One()
	return userContactRequest, ormErr
}

func (model UserContactRequest) Update() db.OrmError {
	_, ormErr := db.Model(model).Update().Update()
	return ormErr
}

func UserContactRequestNew(userId, ContactId uint64, Remark string) *UserContactRequest {
	_userContactRequest := &UserContactRequest{}
	_userContactRequest.UserId = userId
	_userContactRequest.ContactId = ContactId
	_userContactRequest.UserRemark = Remark
	_userContactRequest.CreateTime = uint64(time.Now().Unix())
	_userContactRequest.UpdateTime = uint64(time.Now().Unix())
	_userContactRequest.Status = UserContactRequestStatusRequest
	_ = db.Model(_userContactRequest).Insert().Insert()
	return _userContactRequest
}

const (
	UserContactRequestStatusRequest uint16 = iota //新请求
	UserContactRequestStatusSuccess               // 通过
	UserContactRequestStatusRefuse                // 拒绝
	UserContactRequestStatusIgnore                // 忽略
	UserContactRequestStatusReply                 // 回复
)
