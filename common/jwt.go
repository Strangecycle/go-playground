package common

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/micro/go-micro/v2/logger"
	"go-playground/config"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Uid      uint   `json:"id"`
	jwt.StandardClaims
}

func GenToken(username string, uid uint) string {
	claims := &Claims{
		username,
		uid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 天过期
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		logger.Info("fail to generate token")
	}

	return tokenString
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JwtSecret), nil
	})

	return token, claims, err
}
