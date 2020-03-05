package main

import (
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/yanlong-li/HelloWorld-GO/io/db"
	"github.com/yanlong-li/HelloWorld-GO/io/network/socket"
	"github.com/yanlong-li/HelloWorld-GO/io/network/websocket"
	_ "github.com/yanlong-li/HelloWorldServer/controller"
	_ "github.com/yanlong-li/HelloWorldServer/model"
	_ "github.com/yanlong-li/HelloWorldServer/packetModel"
)

func main() {

	db.ConfigDb("mysql", "root:123456@tcp(127.0.0.1:3306)/hello_world?charset=utf8")

	go websocket.Server(":3001")
	socket.Server(":3000")
}
