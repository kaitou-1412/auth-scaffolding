package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
)

func protected(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {

		// claims, _ = context.Get("claims")

		// Successful access to protected resource
		slog.Info(SuccessProtectedRoute)
		context.JSON(
			http.StatusOK,
			gin.H{
				"status":  http.StatusOK,
				"message": SuccessProtectedRoute,
				"data": map[string]string{
					"resource": "Hello world!",
				},
			})
	}
}
