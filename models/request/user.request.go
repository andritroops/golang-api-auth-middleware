package request

type UserCreateRequest struct {
	Name        string `json:"name" form:"name" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
	Email       string `json:"email" form:"email" validate:"required"`
	Password    string `json:"password" form:"password" validate:"required"`
	Role        string `json:"role_id" form:"role_id"`
	File        string `json:"file" form:"file"`
}

type UserUpdateRequest struct {
	Name        string `json:"name" form:"name" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
	Email       string `json:"email" form:"email" validate:"required"`
}

type UserPasswordRequest struct {
	Password string `json:"password" form:"password" validate:"required"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}
