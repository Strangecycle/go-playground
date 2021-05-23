package main

import (
	"github.com/gin-gonic/gin"
	"go-playground/config"
	"go-playground/gateway/router"
	"log"
)

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	router.CollectRoutes(r)
	if err := r.Run(":" + config.Gateway); err != nil {
		log.Fatal(err.Error())
	}
}
