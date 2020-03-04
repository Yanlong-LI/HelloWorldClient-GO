package model

import "github.com/yanlong-li/HelloWorld-GO/io/db"

type Channel struct {
	Id           uint64
	ServerId     uint64
	Name         string
	Avatar       string
	CreateTime   uint64
	CreateUserId uint64
	OwnerUserId  uint64
	Status       uint8
	ParentId     uint64
	UpdateTime   uint64
	//Region       string
	Describe string
}

func GetChannels() []interface{} {
	return db.Model(&Channel{}).Where("parent_id", 0).All()
}

func (cha *Channel) GetChildren() []interface{} {
	return db.Model(&Channel{}).Where("parent_id", cha.Id).All()
}
