package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email"`
	Sentence string `json:"sentence"`
}
