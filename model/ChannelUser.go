package model

type ChannelUser struct {
	Id         uint64
	ChannelId  uint64
	UserId     uint64
	CreateTime uint64
	UpdateTime uint64
	OpenId     string
}
