package user

import (
	"time"
)

type User struct {
	Id   uint32
	Name string
	Time time.Time
}
