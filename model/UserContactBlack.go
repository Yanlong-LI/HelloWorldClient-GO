package model

type UserContactBlack struct {
	Id         uint64
	UserId     uint64
	ContactId  uint64
	CreateTime uint64
	Remark     string
}

func (uc *UserContactBlack) GetUser() (User, db.OrmError) {
	return GetUserById(uc.UserId)
}

func (uc *UserContactBlack) GetContactUser() (User, db.OrmError) {
	return GetUserById(uc.ContactId)
}
