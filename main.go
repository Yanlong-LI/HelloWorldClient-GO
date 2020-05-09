package main

import (
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/yanlong-li/HelloWorld-GO/io/db"
	"github.com/yanlong-li/HelloWorld-GO/io/network/socket"
	_ "github.com/yanlong-li/HelloWorldServer/controller"
)

func main() {

	db.ConfigDb("mysql", "helloworld:helloworld@tcp(localhost:3306)/helloworld?charset=utf8")

	//go websocket.Server(":3001")
	socket.Server(":3000")
}
