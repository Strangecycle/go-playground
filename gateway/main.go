package main

import (
	"github.com/gin-gonic/gin"
	"go-playground/gateway/handler"
	"log"
)

func main() {
	r := gin.Default()

	// Register client service and get controller
	apiHandler := handler.GetAPIHandler()

	userGroup := r.Group("/user")
	{
		userGroup.GET("/", apiHandler.AddUser)
	}

	if err := r.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
