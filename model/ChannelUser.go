package model

type ChannelUser struct {
	Id         uint64
	ChannelId  uint64
	UserId     uint64
	CreateTime uint64
	UpdateTime uint64
	OpenId     string
	DeleteTime uint64
}

func (cu *ChannelUser) Channel() (Channel, db.OrmError) {
	return getChannelById(cu.ChannelId)
}
func (cu *ChannelUser) User() (User, db.OrmError) {
	return GetUserById(cu.UserId)
}

func GetUserChannels(userId uint64) []interface{} {
	return db.Model(&ChannelUser{}).Find().Where("=", "user_id", userId).AndWhere("delete_time", 0).All()
}
