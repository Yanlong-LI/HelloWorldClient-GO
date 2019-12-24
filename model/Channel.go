package model

import "HelloWorld/io/db"

type Channel struct {
	Id           uint64
	ServerId     uint64
	Name         string
	Icon         string
	CreateTime   uint64
	CreateUserId uint64
	OwnerUserId  uint64
	Status       uint8
	ParentId     uint64
}

func GetChannels() []interface{} {
	return db.Find(&Channel{}).Where("parent_id", 0).All()
}

func (cha *Channel) GetChildren() []interface{} {
	return db.Find(&Channel{}).Where("parent_id", cha.Id).All()
}
