package main

import (
	"github.com/gin-gonic/gin"
	"go-playground/config"
	"go-playground/gateway/router"
	"log"
)

func main() {
	r := gin.Default()
	router.CollectRoutes(r)
	if err := r.Run(":" + config.GATEWAY); err != nil {
		log.Fatal(err.Error())
	}
}
