package middleware

import (
	"e-commerce-go/internal/helpers"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		token := strings.TrimPrefix(auth, "Bearer ")

		claims, err := helpers.VerifyJWT(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			return
		}

		if claims.Type != "access-token" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			return
		}

		if claims.ExpiresAt.Before(time.Now()) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token telah kadaluarsa"})
			return
		}

		// Set UserID ke context
		ctx.Set("userID", claims.UserID)
		ctx.Next()
	}
}
