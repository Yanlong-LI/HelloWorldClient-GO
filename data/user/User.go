package user

import (
	"time"
)

type User struct {
	Id   uint32
	Name string
	Time time.Time
}

var addChan = make(chan User)
var delChan = make(chan uint32)

func init() {

	go func() {

		for {
			select {
			case user := <-addChan:
				List[user.Id] = user
			case id := <-delChan:
				delete(List, id)
			}
		}
	}()
}

var List = make(map[uint32]User)

func Register(user User) {
	addChan <- user
}
func GetUser(ID uint32) User {
	return List[ID]
}

func GetUsers() map[uint32]User {
	return List
}
func RemoveUser(ID uint32) {
	delChan <- ID
}
