package model

type ServerUser struct {
	Id         uint64
	ServerId   uint64
	CreateTime uint64
	UpdateTime uint64
	UserId     uint64
	OpenId     string
}
