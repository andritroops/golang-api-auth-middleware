package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type GetUser struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserCreateRequest struct {
	Name        string    `gorm:"type:varchar(191)" form:"name" validate:"required,min=3,max=32"`
	PhoneNumber string    `gorm:"type:varchar(13)" form:"phone_number" validate:"required,min=9,max=13"`
	Email       string    `gorm:"type:varchar(191)" form:"email" validate:"required,email,min=6,max=32"`
	Password    string    `gorm:"type:varchar(191)" form:"password" validate:"required"`
	CreatedAt   time.Time `gorm:"type:timestamp"`
	UpdatedAt   time.Time `gorm:"type:timestamp"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateStruct(user UserCreateRequest) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
