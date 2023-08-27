package middleware

import (
	"eMenu-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, err := utils.ValidateToken(ctx)
		if err != nil {
			ctx.String(http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return
		}
		ctx.Set("user_id", claims["user_id"])
		ctx.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, err := utils.ValidateToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		role, ok := claims["role"].(string)
		if !ok || role != "admin" {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
