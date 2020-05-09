package model

import "github.com/yanlong-li/HelloWorld-GO/io/db"

type UserContact struct {
	Id         uint64
	UserId     uint64
	ContactId  uint64
	Remarks    string
	CreateTime uint64
	UpdateTime uint64
	Status     uint16 //0 待通过\r\n1 已通过\r\n2 黑名单\r\n4 删除 8 正在等待对方审核
}

func (uc *UserContact) GetUserInfo() (User, db.OrmError) {
	return GetUserById(uc.UserId)
}

func (uc *UserContact) GetContactInfo() (User, db.OrmError) {
	return GetUserById(uc.ContactId)
}
