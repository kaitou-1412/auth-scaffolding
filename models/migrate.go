package models

import (
	"auth/db"
	"gorm.io/gorm"
)

func DBSetup() (*gorm.DB, error) {
	// Initiate the database connection
	DB, err := db.InitDB()
	if err != nil {
		return nil, err
	}

	// Auto-migrates db schema
	if err := AutoMigrateUsers(DB); err != nil {
		return nil, err
	}

	return DB, nil
}
