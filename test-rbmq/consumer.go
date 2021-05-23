package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
	"go-playground/config"
)

func main() {
	consulReg := consul.NewRegistry(registry.Addrs(config.Registry))

	rqBroker := rabbitmq.NewBroker(broker.Addrs("amqp://coven:123456@39.102.51.46:5672/my_vhost"))
	rqBroker.Init()
	rqBroker.Connect()

	service := micro.NewService(
		micro.Name("go.micro.coven.notification"),
		micro.Registry(consulReg),
		micro.Broker(rqBroker),
	)

	service.Init()

	/*
		go micro 用相同的 API 整合了不同的消息中间件（mqtt, rabbitMQ 等），
		主要使用步骤：
			1、连接与注册消息中间件（micro/plugin/broker 下支持的中间件），拿到中间件 broker 实例
			2、消息服务端（consumer）broker.Subscribe('event', callback) 方法订阅（监听）一个消息事件
			3、发布消息端（publisher）重复步骤 1
			4、broker.Publish() 发布一条消息，进入步骤 2 的 callback
		PS: 其它的中间件大致使用流程也差不多，go micro 会将中间件整合到固定 API 中实现可插拔式的使用模式
	*/

	sub, _ := rqBroker.Subscribe("message", func(event broker.Event) error {
		fmt.Println("收到客户端发来的消息：", string(event.Message().Body))
		fmt.Println("发送消息的人是：", event.Message().Header["name"])
		return nil
	})
	sub2, _ := rqBroker.Subscribe("sendCaptcha", func(event broker.Event) error {
		fmt.Println("收到消息，开始发送验证码：", string(event.Message().Body))
		return nil
	})
	defer sub.Unsubscribe()
	defer sub2.Unsubscribe()

	if err := service.Run(); err != nil {
		logger.Fatal(err.Error())
	}
}
