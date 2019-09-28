package user

import (
	"time"
)

type User struct {
	Name string
	Time time.Time
}

var List = make(map[uint32]User)

func Register(ID uint32, user User) {
	List[ID] = user
}
func GetUser(ID uint32) User {
	return List[ID]
}

func GetUsers() map[uint32]User {
	return List
}
func RemoveUser(ID uint32) {
	delete(List, ID)
}
