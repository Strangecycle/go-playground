package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	client "github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/util/log"
	"go-playground/proto/user"
	"net/http"
	"time"
)

func (h *APIHandler) AddUser(ctx *gin.Context) {
	request := user.AddUserRequest{
		Phone: "18100751803",
	}
	response, err := h.userClient.AddUser(context.TODO(), &request, func(options *client.CallOptions) {
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
