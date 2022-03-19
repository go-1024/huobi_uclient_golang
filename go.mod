module github.com/gostudys/huobi_uclient_golang

go 1.16

require (
	github.com/gorilla/websocket v1.5.0
	go.uber.org/zap v1.21.0
)
//require (
//	github.com/gostudys/huobi_uclient_golang/sdk/linearswap v0.0.0
//)
replace "github.com/gostudys/huobi_uclient_golang/sdk/linearswap" => "../sdk/linearswap"
