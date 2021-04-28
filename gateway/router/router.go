package router

import (
	"github.com/gin-gonic/gin"
	"go-playground/gateway/middleware"
)

func CollectRoutes(r *gin.Engine) {
	r.Use(middleware.CorsMiddleware())
	CreateUserRouter(r)
}
