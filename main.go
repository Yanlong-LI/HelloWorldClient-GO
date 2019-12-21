package main

import (
	"HelloWorld/io/db"
	_ "HelloWorldServer/controller"
	_ "HelloWorldServer/model"
	"HelloWorldServer/model/user"
	_ "HelloWorldServer/packet"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func main() {

	db.ConfigDb("mysql", "root:123456@tcp(127.0.0.1:3306)/hello_world?charset=utf8")

	//c:=db.Find(&user.User{}).Where([]interface{}{"and", []interface{}{"id", "1"}, []interface{}{"name", "王二狗"}, []interface{}{"age", 18}}).AndWhere([]interface{}{"and", []interface{}{"id", "1"}, []interface{}{"name", "王二狗"}, []interface{}{"age", 18}}).OrWhere([]interface{}{"and", []interface{}{"id", "1"}, []interface{}{"name", "王二狗"}, []interface{}{"age", 18}}).Sql()
	c := db.Find(&user.User{}).Where([]interface{}{"and", []interface{}{"or", []interface{}{"id", "1"}, []interface{}{"id", "3"}, []interface{}{"id", "2"}}, []interface{}{">", "id", "0"}, []interface{}{">=", "id", "1"}}).All()
	fmt.Println(c)

	//go websocket.Server(":3001")
	//socket.Server(":3000")
}
