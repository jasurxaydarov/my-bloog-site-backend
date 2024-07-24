package middlewares

import (
	"jasurxaydarov/my-bloog-site-backend/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("authorization")

		if tokenString == "" {
			ctx.JSON(401, gin.H{"error": "authorization token not provided"})
			ctx.Abort()
		}

		claim, err := token.ParseJWT(tokenString)
		if err != nil {
			ctx.JSON(401, gin.H{"error": err.Error()})
			ctx.Abort()
		}

		if claim.UserRole != "viwer" {
			ctx.JSON(401, gin.H{"error": "your role isn't viwer "})
			ctx.Abort()
		}

		ctx.Next()
	}
}
