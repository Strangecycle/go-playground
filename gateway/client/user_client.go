package client

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"go-playground/config"
	"go-playground/proto/user"
)

var userClient user.UserService

func GetUserClient() user.UserService {
	if userClient != nil {
		return userClient
	}

	consulReg := consul.NewRegistry(registry.Addrs(config.Registry))

	service := micro.NewService(
		micro.Registry(consulReg),
	)

	return user.NewUserService("go.micro.coven.user", service.Client())
}
