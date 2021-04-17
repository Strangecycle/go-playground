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
	*response = service.UserLogin(request.Phone, request.Captcha)
	return nil
}

func Handler() user.UserHandler {
	return new(handler)
}
