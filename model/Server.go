package model

type Server struct {
	Id         uint64
	Name       string
	Version    string
	CreateTime uint64
	Region     string
	UpdateTime uint64
	DeleteTime uint64
	Status     byte
}
