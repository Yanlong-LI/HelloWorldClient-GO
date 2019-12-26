package model

type UserContact struct {
	Id         uint64
	UserId     uint64
	ContactId  uint64
	Remarks    string
	CreateTime uint64
	UpdateTime uint64
}

func (uc *UserContact) GetUserInfo() (User, error) {
	return GetUserById(uc.UserId)
}

func (uc *UserContact) GetContactInfo() (User, error) {
	return GetUserById(uc.ContactId)
}
