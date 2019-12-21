package user

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(&User{})
}

type User struct {
	Id       uint64 `db:"id"`
	Nickname string `db:"nickname"`
}
