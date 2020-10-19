package main

import (
	_ "github.com/go-sql-driver/mysql" // import your used driver
	_ "github.com/yanlong-li/hi-go-server/controller"
	"github.com/yanlong-li/hi-go-socket/socket"
	"github.com/yanlong-li/hi-go-socket/websocket"
)

func main() {

	db.ConfigDb("mysql", "helloworld:helloworld@tcp(localhost:3306)/helloworld?charset=utf8")

	go websocket.Server(":3001")
	socket.Server(":3000")
}
