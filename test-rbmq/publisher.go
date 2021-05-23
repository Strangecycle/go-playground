package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"
)

func main() {
	rqBroker := rabbitmq.NewBroker(broker.Addrs("amqp://coven:123456@39.102.51.46:5672/my_vhost"))
	rqBroker.Init()
	rqBroker.Connect()

	err := rqBroker.Publish("message", &broker.Message{
		Header: map[string]string{
			"name": "Coven",
		},
		Body: []byte("这是客户端的消息体"),
	})
	if err != nil {
		logger.Error(err.Error())
		return
	}

	err2 := rqBroker.Publish("sendCaptcha", &broker.Message{
		Body: []byte("111000"),
	})
	if err2 != nil {
		logger.Error(err.Error())
	}
	fmt.Println("发送消息成功")
}
