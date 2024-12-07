package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log/slog"
)

type User struct {
	ModelUUID
	Name     string   `json:"name" binding:"required"`
	Username string   `json:"username" binding:"required" gorm:"unique"`
	Email    string   `json:"email" binding:"required" gorm:"unique"`
	Password string   `json:"password"`
	Role     RoleType `json:"role" gorm:"default:'user'; type:varchar(20); check:role IN ('user', 'manager', 'admin')"`
	Active   bool     `json:"active" gorm:"default:true"`
}

// AutoMigrateUsers createTables creates the "users" table in the database if it doesn't exist.
func AutoMigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	if err != nil {
		slog.Error("Could not create users table.")
		return err
	}
	return nil
}

func (u *User) UsernameExists(db *gorm.DB) bool {
	var user User
	result := db.Where(&User{Username: u.Username}).First(&user)
	return result.RowsAffected > 0
}

func (u *User) EmailExists(db *gorm.DB) bool {
	var user User
	result := db.Where(&User{Email: u.Email}).First(&user)
	return result.RowsAffected > 0
}

func (u *User) HashPassword() error {
	// Default cost is 10, higher cost = more secure but slower
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		slog.Error("Could not hash password.")
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) GetUserByUsername(db *gorm.DB) {
	db.Where(&User{Username: u.Username}).First(u)
}
