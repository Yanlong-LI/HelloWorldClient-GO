package model

type UserToken struct {
	Id         uint64
	UserId     uint64
	Token      string
	ExpireTime uint64
	CreateTime uint64
	UpdateTime uint64
}

func (ut *UserToken) User() (User, db.OrmError) {
	return GetUserById(ut.UserId)
}

func GetUserByToken(Token string) (User, db.OrmError) {
	var userToken = UserToken{}
	err := db.Model(&userToken).Find().Where("token", Token).One()
	if err.Empty() {
		return User{}, err
	}
	return GetUserById(userToken.UserId)
}
