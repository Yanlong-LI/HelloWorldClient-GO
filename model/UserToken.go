package model

import "github.com/yanlong-li/HelloWorld-GO/io/db"

type UserToken struct {
	Id         uint64
	UserId     uint64
	Token      string
	ExpireTime uint64
	CreateTime uint64
	UpdateTime uint64
}

func (ut *UserToken) User() (User, error) {
	return GetUserById(ut.UserId)
}

func GetUserByToken(Token string) (User, error) {
	var userToken = UserToken{}
	err := db.Model(&userToken).Where("token", Token).One()
	if err != nil {
		return User{}, err
	}
	return GetUserById(userToken.UserId)
}
