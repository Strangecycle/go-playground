package handler

import (
	"context"
	"go-playground/proto/user"
	"go-playground/user-service/service"
)

type handler struct {
}

func (h handler) SendCaptcha(ctx context.Context, request *user.CaptchaRequest, response *user.CaptchaResponse) error {
	*response = service.SendCaptcha(request)
	return nil
}

func (h handler) UserLogin(ctx context.Context, request *user.UserLoginRequest, response *user.UserLoginResponse) error {
	*response = service.UserLogin(request.GetPhone(), request.GetCaptcha())
	return nil
}

func (h handler) UserInfo(ctx context.Context, request *user.UserInfoRequest, response *user.UserInfoResponse) error {
	*response = service.UserInfo(request.GetPhone(), request.GetEmail())
	return nil
}

func (h handler) UserEdit(ctx context.Context, request *user.UserEditRequest, response *user.UserEditResponse) error {
	*response = service.UserEdit(request)
	return nil
}

func Handler() user.UserHandler {
	return new(handler)
}
