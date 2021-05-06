package client

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"go-playground/config"
	"go-playground/proto/file"
)

var fileService file.FileService

func GetFileService() file.FileService {
	if fileService != nil {
		return fileService
	}

	consulReg := consul.NewRegistry(
		registry.Addrs(config.REGISTRY),
	)

	service := micro.NewService(
		micro.Registry(consulReg),
	)

	return file.NewFileService("go.micro.coven.file", service.Client())
}
