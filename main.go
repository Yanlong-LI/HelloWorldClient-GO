package main

import (
	"HelloWorld/io/db"
	"HelloWorld/io/network/socket"
	"HelloWorld/io/network/websocket"
	_ "HelloWorldServer/controller"
	_ "HelloWorldServer/model"
	_ "HelloWorldServer/packet"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func main() {

	db.ConfigDb("mysql", "root:123456@tcp(127.0.0.1:3306)/hello_world?charset=utf8")

	go websocket.Server(":3001")
	socket.Server(":3000")
}
