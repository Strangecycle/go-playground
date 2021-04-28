package router

import (
	"github.com/gin-gonic/gin"
	"go-playground/gateway/handler"
	"go-playground/gateway/middleware"
)

func CreateUserRouter(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userHandler := handler.GetUserHandler()
		userGroup.GET("/send", userHandler.Send)
		userGroup.POST("/login", userHandler.Login)
		userGroup.GET("/info", middleware.AuthMiddleware(), userHandler.Info)
	}
}
