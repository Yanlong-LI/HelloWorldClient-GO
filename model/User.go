package model

import (
	"github.com/yanlong-li/HelloWorld-GO/io/db"
)

type User struct {
	Id         uint64
	Nickname   string
	CreateTime uint64
	Avatar     string
	Language   string
	Region     string
	UpdateTime uint64
}

func GetUserById(Id uint64) (User, db.OrmError) {
	var _user User
	err := db.Model(&_user).Find().Where("=", "id", Id).One()
	return _user, err
}
