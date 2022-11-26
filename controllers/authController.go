package controllers

import (
	"log"
	"time"

	"github.com/andritroops/go-latihan/config"
	"github.com/andritroops/go-latihan/models/entity"
	"github.com/andritroops/go-latihan/models/request"
	"github.com/andritroops/go-latihan/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Signin(ctx *fiber.Ctx) error {

	signinRequest := new(request.SigninRequest)

	if err := ctx.BodyParser(signinRequest); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})

	}

	var errors []*request.ErrorResponse
	validate := validator.New()
	errValidate := validate.Struct(signinRequest)

	if errValidate != nil {

		for _, err := range errValidate.(validator.ValidationErrors) {
			var element request.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Field()
			errors = append(errors, &element)
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	var user entity.User

	// Chek Validation Email
	err := config.DB.First(&user, "email = ?", signinRequest.Email).Error

	if err != nil {

		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})

	}

	// Check Validation Password
	isValid := utils.CheckPasswordHash(signinRequest.Password, user.Password)

	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	// Generate JWT
	claims := jwt.MapClaims{}

	claims["name"] = user.Name
	claims["phone_number"] = user.PhoneNumber
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	claims["role"] = user.Role

	token, errGenerateToken := utils.GenerateToken(&claims)

	if errGenerateToken != nil {
		log.Println(errGenerateToken)

		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized ",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})

}
