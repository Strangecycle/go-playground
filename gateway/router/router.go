package router

import "github.com/gin-gonic/gin"

func CollectRoutes(r *gin.Engine) {
	CreateUserRouter(r)
}
