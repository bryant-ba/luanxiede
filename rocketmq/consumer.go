package main

import (
	"fmt"
	"gitee.com/zhucheer/orange/queue"
	"time"
)

func main() {
	var mqip string
	fmt.Println("请输入MQnameserverIP:")
	fmt.Scanln(&mqip)
	mqConsumerClient := queue.RegisterRocketConsumerMust([]string{mqip + ":9876"}, "test")
	mqConsumerClient.ListenReceiveMsgDo("topicTest", func(mqMsg queue.MqMsg) {
		// 收到一条消息
		fmt.Println("receive====>", mqMsg.MsgId, mqMsg.BodyString())

	})
	time.Sleep(20 * time.Second)
}
