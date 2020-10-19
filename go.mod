module github.com/yanlong-li/hi-go-server

go 1.14

require (
	github.com/ProtonMail/gopenpgp/v2 v2.0.0
	github.com/go-sql-driver/mysql v1.5.1-0.20200818111213-46351a889297
	github.com/yanlong-li/hi-go-logger v0.0.0-20201019104050-b1e94d395fee
	github.com/yanlong-li/hi-go-orm v0.0.0-20201019094537-a7dd8b139729
	github.com/yanlong-li/hi-go-socket v0.0.0-20201019105643-c29816f01818
)

replace (
	github.com/yanlong-li/hi-go-logger => ../hi-go-logger
	github.com/yanlong-li/hi-go-orm => ../hi-go-orm
	github.com/yanlong-li/hi-go-socket => ../hi-go-socket
	golang.org/x/crypto => github.com/ProtonMail/crypto v0.0.0-20201016191319-576ad9c42ffa
)
