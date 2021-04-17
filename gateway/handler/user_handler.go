package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/logger"
	c "go-playground/gateway/client"
	"go-playground/gateway/response"
	"go-playground/proto/user"
)

/*
	Go 的 interface 是隐式实现
	只要 struct 实现了 interface 里的方法，那么这个 struct 就实现了这个 interface

	注意：如果是指针接收者实现了接口，那么接口中就只能存储指针类型，如
		func (u *UserHandler) CreateUser() {}
		var iu IUserHandler
		iu = UserHandler{} // 不可以，因为接收者是指针类型实现
		iu = &UserHandler{} // 可以，Go 内部会自动 * 取值
		如果是值类型则二者都能存储
*/

type IUserHandler interface {
	Send(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type UserHandler struct {
	userClient user.UserService
}

func GetUserHandler() IUserHandler {
	return UserHandler{
		userClient: c.GetUserClient(),
	}
}

func (uh UserHandler) Send(ctx *gin.Context) {
	var request user.CaptchaRequest
	captcha, err := uh.userClient.SendCaptcha(context.Background(), &request)
	if err != nil {
		logger.Info(err.Error())
		response.ServerError(ctx)
		return
	}
	response.Success(ctx, captcha.GetCaptcha())
}

func (uh UserHandler) Login(ctx *gin.Context) {
	request := user.UserLoginRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		logger.Info(err.Error())
		response.Fail(ctx)
		return
	}

	loginResponse, err := uh.userClient.UserLogin(context.Background(), &request)
	if err != nil {
		logger.Info(err.Error())
		response.ServerError(ctx)
		return
	}

	// 验证码错误
	if loginResponse.GetToken() == "" {
		response.Fail(ctx)
		return
	}

	response.Success(ctx, loginResponse.GetToken())
}
