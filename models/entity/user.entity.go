package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	PhoneNumber string         `json:"phone_number"`
	Email       string         `json:"email"`
	Role        string         `json:"role_id"`
	Password    string         `json:"-"`
	File        string         `json:"file"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index, column:deleted_at"`
}
