package request

import "github.com/go-playground/validator/v10"

type UserCreateRequest struct {
	Name        string `json:"name" form:"name" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
	Email       string `json:"email" form:"email" validate:"required"`
	Password    string `json:"password" form:"password" validate:"required"`
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
			element.Value = err.Field()
			errors = append(errors, &element)
		}
	}

	return errors
}
