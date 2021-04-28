package middleware

import (
	"github.com/gin-gonic/gin"
	"go-playground/common"
	"go-playground/gateway/response"
	"go-playground/user-service/model"
	"strings"
)

// AuthMiddleware for validate token
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		db := common.GetDB()
		var user model.User

		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			unauthorized(ctx)
			return
		}

		tokenMain := token[7:]
		parseToken, claims, err := common.ParseToken(tokenMain)
		if err != nil || !parseToken.Valid {
			unauthorized(ctx)
			return
		}

		userId := claims.Uid
		db.First(&user, userId)

		if user.ID == 0 {
			unauthorized(ctx)
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}

func unauthorized(ctx *gin.Context) {
	response.Unauthorized(ctx)
	ctx.Abort()
}
