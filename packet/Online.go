package packet

//新用户上线
type UserOnline struct {
	User
}

type User struct {
	Id       uint64 //用户ID
	UserName string //用户昵称
	Avatar   string //用户头像
	//RegisterTime uint64 //注册时间
}
type UserList struct {
	UserList []User
}
