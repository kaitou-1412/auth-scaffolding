package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ModelUUID struct {
	ID        uuid.UUID      `json:"id" gorm:"primaryKey; type:uuid; default:gen_random_uuid()"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
