package model

// 用户密码表
type UserPassword struct {
	Id         uint64
	UserId     uint64
	Type       uint8 // 0登陆密码 1 支付密码
	Password   string
	CreateTime uint64
	UpdateTime uint64
}
