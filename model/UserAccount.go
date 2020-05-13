package model

import "github.com/yanlong-li/HelloWorld-GO/io/db"

// 用户账户表
type UserAccount struct {
	Id         uint64
	UserId     uint64
	Account    string
	Type       uint8 // 0 邮箱 1 手机号码
	CreateTime uint64
	UpdateTime uint64
}

func (uc *UserAccount) GetUser() (User, db.OrmError) {
	return GetUserById(uc.UserId)
}
