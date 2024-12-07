package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(server *gin.Engine, db *gorm.DB) {
	server.POST("/api/auth/signup", signup(db))
	server.POST("/api/auth/login", login(db))
	server.POST("/api/auth/refresh", refresh(db))
	server.GET("/api/protected", authMiddleware(), protected(db))
}
