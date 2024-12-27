package utils

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

var (
	DbHost     = ""
	DbUser     = ""
	DbPassword = ""
	DbName     = ""
	DbPort     = ""
)

func LoadEnvVars() error {
	env := os.Getenv("ENV")
	if env == "" {
		env = "production"
	}

	// Load .env file only in development
	if env == "development" {
		slog.Info("Loading .env file for development environment...")
		err := godotenv.Load()
		if err != nil {
			slog.Error("Error loading .env file")
			return err
		}
	}

	DbHost = os.Getenv("POSTGRES_HOST")
	DbUser = os.Getenv("POSTGRES_USER")
	DbPassword = os.Getenv("POSTGRES_PASSWORD")
	DbName = os.Getenv("POSTGRES_DB")
	DbPort = os.Getenv("POSTGRES_PORT")

	return nil

}
