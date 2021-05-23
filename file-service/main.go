package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"go-playground/config"
	"go-playground/file-service/handler"
	"go-playground/proto/file"
	"time"
)

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs(config.Registry),
	)

	service := micro.NewService(
		micro.Name("go.micro.coven.file"),
		micro.Registry(consulReg),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
	)

	service.Init()

	file.RegisterFileHandler(service.Server(), handler.Handler())

	if err := service.Run(); err != nil {
		logger.Fatal(err.Error())
	}
}
