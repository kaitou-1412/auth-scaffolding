package main

import (
	"auth/models"
	"auth/routes"
	"auth/utils"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
		// Load environment variables
	if err := utils.LoadEnvVars(); err != nil {
		os.Exit(1)
	}

	// DB Setup
	db, err := models.DBSetup()
	if err != nil {
		os.Exit(1)
	}

	// Create a new Gin server instance with default middleware
	server := gin.Default()

	// Register the routes for the server
	routes.RegisterRoutes(server, db)

	// Run the server on port 8080
	err = server.Run(":8080")
	if err != nil {
		slog.Error("Could not start server", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
