package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	client "github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/util/log"
	c "go-playground/gateway/client"
	"go-playground/proto/user"
	"net/http"
	"time"
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
	AddUser(ctx *gin.Context)
}

type UserHandler struct {
	userClient user.UserService
}

func GetUserHandler() IUserHandler {
	return UserHandler{
		userClient: c.GetUserClient(),
	}
}

func (uh UserHandler) AddUser(ctx *gin.Context) {
	request := user.AddUserRequest{
		Phone: "18100751803",
	}
	response, err := uh.userClient.AddUser(context.Background(), &request, func(options *client.CallOptions) {
		// 设置调用服务超时时间，默认 5s
		// 如果超时会报 408 Request Timeout 错误
		options.RequestTimeout = time.Second * 30
		options.DialTimeout = time.Second * 30
	})
	if err != nil {
		log.Info(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": response.GetMessage(),
	})
}
