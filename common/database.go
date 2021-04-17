package common

import (
	"go-playground/config"
	"go-playground/user-service/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func init() {
	open, err := gorm.Open(mysql.New(mysql.Config{DSN: config.DSN}), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail to connect database.")
	}
	open.AutoMigrate(&model.User{})
	db = open
}

func GetDB() *gorm.DB {
	return db
}
