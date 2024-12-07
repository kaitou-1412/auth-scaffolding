package utils

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"
)

var (
	DbHost     = ""
	DbUser     = ""
	DbPassword = ""
	DbName     = ""
	DbPort     = ""
)

func LoadEnvVars() error {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
		return err
	}

	DbHost = os.Getenv("POSTGRES_HOST")
	DbUser = os.Getenv("POSTGRES_USER")
	DbPassword = os.Getenv("POSTGRES_PASSWORD")
	DbName = os.Getenv("POSTGRES_DB")
	DbPort = os.Getenv("POSTGRES_PORT")

	return nil

}
