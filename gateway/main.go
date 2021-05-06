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
	// 开放静态资源，当访问 /upload 时将开放 ./upload 目录下的文件
	r.Static("/upload", "./upload")
	router.CollectRoutes(r)
	if err := r.Run(":" + config.GATEWAY); err != nil {
		log.Fatal(err.Error())
	}
}
