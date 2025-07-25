package middleware

import (
	"e-commerce-go/internal/helpers"
	"e-commerce-go/internal/models"
	"e-commerce-go/pkg"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(roles ...any) gin.HandlerFunc {
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

		var pat models.User
		err = pkg.DB.Where("id = ?", claims.UserID).First(&pat).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return 
		}

		// engga di cek ke tabel personal acces token untuk mempermudah query aja, 
		// karna expired nya untuk acces-token sama dengan yang ada di jtw

		userRole, _ := pat.Role.Value()
		for _, r := range roles {
			if r == userRole {
				ctx.Set("user", pat)
				ctx.Next()
				return
			}
		}
		messege := fmt.Sprintf("Access denied, you don't have access to this resource, your role is %s", userRole)

		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": messege})
	}
}
