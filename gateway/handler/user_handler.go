package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/logger"
	c "go-playground/gateway/client"
	"go-playground/gateway/response"
	"go-playground/proto/user"
	"go-playground/user-service/model"
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
	Info(ctx *gin.Context)
}

type UserHandler struct {
	userClient user.UserService
}

func GetUserHandler() IUserHandler {
	return UserHandler{
		userClient: c.GetUserClient(),
	}
}

// Send
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

// Login
func (uh UserHandler) Login(ctx *gin.Context) {
	request := user.UserLoginRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		logger.Info(err.Error())
		response.Fail(ctx)
		return
	}

	loginResponse, err := uh.userClient.UserLogin(context.Background(), &request)
	// 服务错误
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

// Info
func (uh UserHandler) Info(ctx *gin.Context) {
	var request user.UserInfoRequest
	phone := ctx.Query("phone")
	email := ctx.Query("email")

	if phone == "" && email == "" {
		// 两个参数都未传，返回当前用户信息
		userCtx, _ := ctx.Get("user")
		currentUser := userCtx.(model.User)
		response.Success(ctx, user.UserInfoResponse{
			Id:        uint64(currentUser.ID),
			Username:  currentUser.Username,
			Avatar:    currentUser.Avatar,
			Phone:     currentUser.Phone,
			Email:     currentUser.Email,
			Sentence:  currentUser.Sentence,
			CreatedAt: currentUser.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: currentUser.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
		return
	} else {
		// 传了手机号或邮箱，通过它们来查找用户
		request.Phone = phone
		request.Email = email
	}

	infoResponse, err := uh.userClient.UserInfo(context.Background(), &request)
	if err != nil {
		logger.Info(err.Error())
		response.ServerError(ctx)
		return
	}

	response.Success(ctx, infoResponse)
}
