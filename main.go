package main

import (
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/yanlong-li/HelloWorld-GO/io/db"
	"github.com/yanlong-li/HelloWorld-GO/io/network/socket"
	_ "github.com/yanlong-li/HelloWorldServer/controller"
)

func main() {

	db.ConfigDb("mysql", "root:339d2cab665a696a@tcp(127.0.0.1:3306)/hello_world?charset=utf8")

	//go websocket.Server(":3001")
	socket.Server(":3000")
}
