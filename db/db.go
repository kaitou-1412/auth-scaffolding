package db

import (
	"auth/utils"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log/slog"
	"time"
)

func InitDB() (*gorm.DB, error) {

	var DB *gorm.DB

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		utils.DbHost,
		utils.DbUser,
		utils.DbPassword,
		utils.DbName,
		utils.DbPort,
	)

	DB, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,  // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		slog.Error("Could not connect to database.")
		return nil, err
	}

	sqlDB, err := DB.DB()
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	if sqlDB == nil {
		err = errors.New("database connection is nil")
		slog.Error(err.Error())
		return nil, err
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return DB, nil
}
