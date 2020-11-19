module github.com/yanlong-li/hi-go-server

go 1.15

require (
	github.com/yanlong-li/hi-go-logger v0.0.0-20201019104050-b1e94d395fee
	github.com/yanlong-li/hi-go-socket v0.0.0-20201019105643-c29816f01818
	github.com/yanlong-li/hi-go-gateway v0.0.0-20201119075128-0a84a9b658ce
)

replace (
	github.com/yanlong-li/hi-go-logger => ../hi-go-logger
	github.com/yanlong-li/hi-go-gateway => ../hi-go-gateway
	github.com/yanlong-li/hi-go-socket => ../hi-go-socket
)
