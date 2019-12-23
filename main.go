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

	//c:=db.Find(&user.User{}).Where([]interface{}{"and", []interface{}{"id", "1"}, []interface{}{"name", "王二狗"}, []interface{}{"age", 18}}).AndWhere([]interface{}{"and", []interface{}{"id", "1"}, []interface{}{"name", "王二狗"}, []interface{}{"age", 18}}).OrWhere([]interface{}{"and", []interface{}{"id", "1"}, []interface{}{"name", "王二狗"}, []interface{}{"age", 18}}).Sql()
	//c := db.Find(&model.User{}).Where([]interface{}{"and", []interface{}{"or", []interface{}{"id", "1"}, []interface{}{"id", "3"}, []interface{}{"id", "2"}}, []interface{}{">", "id", "0"}, []interface{}{">=", "id", "1"}}).Limit(0,3).All()
	//fmt.Println(c)

	go websocket.Server(":3001")
	socket.Server(":3000")
}
