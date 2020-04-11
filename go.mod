module github.com/yanlong-li/HelloWorldServer

go 1.14

require (
	github.com/ProtonMail/gopenpgp/v2 v2.0.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/yanlong-li/HelloWorld-GO v0.0.0-20200308072225-efccc577e911
)

replace (
    golang.org/x/crypto => github.com/ProtonMail/crypto v0.0.0-20191122234321-e77a1f03baa0
	github.com/yanlong-li/HelloWorld-GO => ../HelloWorld-GO
)
