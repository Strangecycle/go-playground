package handler

import (
	"go-playground/gateway/client"
	"go-playground/proto/user"
)

type APIHandler struct {
	userClient user.UserService
}

func GetAPIHandler() APIHandler {
	return APIHandler{
		client.GetUserClient(),
	}
}
