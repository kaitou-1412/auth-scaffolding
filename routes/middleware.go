package routes

import (
	"auth/jwt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func authMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Verify Access Token
		token := context.GetHeader("Authorization")
		claims, ok := jwt.VerifyTokenAndGetClaims(token, jwt.AccessToken)
		if !ok {
			slog.Error(ErrProtectedRoute["Authorization"])
			context.JSON(
				http.StatusUnauthorized,
				gin.H{
					"status":  http.StatusUnauthorized,
					"message": ErrProtectedRoute["Authorization"],
				})
			context.Abort()
			return
		}

		// Pass the claims along to the request context
		context.Set("claims", claims)

		// Proceed with the request
		context.Next()
	}
}
