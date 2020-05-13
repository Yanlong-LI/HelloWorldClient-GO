package model

import (
	"github.com/yanlong-li/HelloWorld-GO/io/db"
	"time"
)

type UserContact struct {
	Id         uint64
	UserId     uint64
	ContactId  uint64
	Remark     string
	CreateTime uint64
	UpdateTime uint64
}

func UserContactAddUser(userId1, userId2 uint64) {

	_contactUser := &UserContact{}
	ormErr := db.Model(_contactUser).Find().Where(map[interface{}]interface{}{"user_id": userId1, "contact_id": userId2}).One()

	if (ormErr).Empty() {
		_contactUser.UserId = userId1
		_contactUser.ContactId = userId2
		_contactUser.CreateTime = uint64(time.Now().Unix())
		_contactUser.UpdateTime = _contactUser.CreateTime
		_ = db.Model(_contactUser).Insert().Insert()
	}
}

func (uc *UserContact) GetUser() (User, db.OrmError) {
	return GetUserById(uc.UserId)
}

func (uc *UserContact) GetContactUser() (User, db.OrmError) {
	return GetUserById(uc.ContactId)
}
