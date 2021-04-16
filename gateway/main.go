package main

import (
	"github.com/gin-gonic/gin"
	"go-playground/gateway/router"
	"log"
)

func main() {
	r := gin.Default()
	router.CollectRoutes(r)
	if err := r.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
