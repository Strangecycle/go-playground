package handler

import (
	"context"
	"go-playground/proto/user"
	"go-playground/user-service/service"
)

type handler struct {
}

func (h handler) AddUser(ctx context.Context, request *user.AddUserRequest, response *user.AddUserResponse) error {
	*response = service.AddUser(request.Phone)
	return nil
}

func (h handler) FindUser(ctx context.Context, request *user.FindUserRequest, response *user.FindUserResponse) error {
	return nil
}

func Handler() user.UserHandler {
	return handler{}
}
