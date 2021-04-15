package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"go-playground/proto/user"
	"go-playground/user-service/handler"
	"log"
)

func main() {
	consulRegistry := consul.NewRegistry(
		registry.Addrs("39.102.51.46:8500"),
	)

	service := micro.NewService(
		micro.Name("go.micro.coven.user"),
		micro.Registry(consulRegistry),
	)

	service.Init()

	user.RegisterUserHandler(service.Server(), handler.Handler())

	if err := service.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
