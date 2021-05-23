package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Format(ctx *gin.Context, httpStatus int, code int, data interface{}, msg string) {
	ctx.JSON(
		httpStatus,
		gin.H{
			"data":     data,
			"err_code": code,
			"err_msg":  msg,
		},
	)
}

func Success(ctx *gin.Context, data interface{}) {
	Format(ctx, http.StatusOK, 0, data, "success")
}

// Fail 表示参数有误
func Fail(ctx *gin.Context) {
	Format(ctx, http.StatusOK, http.StatusBadRequest, nil, "bad request")
}

// ServerError 表示服务调用失败
func ServerError(ctx *gin.Context) {
	Format(ctx, http.StatusInternalServerError, http.StatusInternalServerError, nil, "consumer error")
}

// Unauthorized 表示未登录
func Unauthorized(ctx *gin.Context) {
	Format(ctx, http.StatusUnauthorized, http.StatusUnauthorized, nil, "unauthorized")
}
