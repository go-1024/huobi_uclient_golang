package main

import (
	"encoding/json"
	"fmt"
	"github.com/gostudys/huobi_uclient_golang/sdk/linearswap/ws"
	"github.com/gostudys/huobi_uclient_golang/sdk/linearswap/ws/response/index"
	"strings"
	"time"
)

func main() {
	type topics struct {
		Topic string `json:"topic"` // K线图代码
	}
	var Topics []topics
	// 收集所有要订阅处理的 topic
	Topics = append(Topics,
		// BTC-USD
		// 1分钟k线图
		topics{
			Topic: "market.BTC-USD.index.1min",
		},
		// 5分钟k线图
		topics{
			Topic: "market.BTC-USD.index.5min",
		},
		// 15分钟k线图
		topics{
			Topic: "market.BTC-USD.index.15min",
		},
		// ETH-USD
		// 1分钟k线图
		topics{
			Topic: "market.ETH-USD.index.1min",
		},
		// 5分钟k线图
		topics{
			Topic: "market.ETH-USD.index.5min",
		},
		// 15分钟k线图
		topics{
			Topic: "market.ETH-USD.index.15min",
		})
	start := "ok"
	for {
		if start == "ok" {
			for _, obj := range Topics {
				go SocketHuoBi(obj.Topic)
			}
			start = "no"
		}
		time.Sleep(time.Second)
	}
}

// 链接火币网

func SocketHuoBi(parentTopic string) {
	Topic := strings.Split(parentTopic, ".")
	// websocket api
	wsmkClient := new(ws.WSIndexClient).Init("")
	wsmkClient.SubIndexKLine(Topic[1], Topic[3], func(data *index.SubIndexKLineResponse) {
		// 火币网推送的回来的数据转换为字符串
		msgData, MarshalJSONErr := json.Marshal(data) //转换成JSON返回的是byte[]
		if MarshalJSONErr != nil {
			fmt.Println(MarshalJSONErr)
			return
		}
		//logger.Info(fmt.Sprintf(
		//	"SUCCESS 当前订阅的 topic 为：%v,返回的数据为：%v",
		//	parentTopic, string(msgData)))
		if len(msgData) > 0 {
			fmt.Println(string(msgData))
			// 收集数据缓存到redis
		} else {
			return
		}

	}, "")
}
