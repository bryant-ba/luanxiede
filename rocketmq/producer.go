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
	mqProducerClient := queue.RegisterRocketProducerMust([]string{mqip + ":9876"}, "test", 1)
	go func() {
		for i := 0; i < 10; i++ {
			// 向队列发送一条消息 填入消息队列topic和消息体信息
			ret, _ := mqProducerClient.SendMsg("topicTest", "Hello mq~~")
			fmt.Println("========producer push one message====", ret.MsgId)
			time.Sleep(time.Second)
		}
	}()
}
