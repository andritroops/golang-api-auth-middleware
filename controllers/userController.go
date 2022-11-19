package controllers

import (
	"log"

	"github.com/andritroops/go-latihan/config"
	"github.com/andritroops/go-latihan/models/entity"
	"github.com/andritroops/go-latihan/models/request"
	"github.com/andritroops/go-latihan/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserIndex(ctx *fiber.Ctx) error {

	var users []entity.User

	result := config.DB.Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.Status(fiber.StatusOK).JSON(users)
}

func UserStore(ctx *fiber.Ctx) error {

	user := new(request.UserCreateRequest)

	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})

	}

	var errors []*request.ErrorResponse
	validate := validator.New()
	errValidate := validate.Struct(user)

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

	newUser := entity.User{
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	newUser.Password = hashedPassword
	errCreateUser := config.DB.Create(&newUser).Error

	if errCreateUser != nil {

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errCreateUser,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    newUser,
	})

}

func UserEdit(ctx *fiber.Ctx) error {

	userID := ctx.Params("id")

	var user entity.User

	err := config.DB.First(&user, "id = ?", userID).Error

	if err != nil {

		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})

	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})

}

func UserUpdate(ctx *fiber.Ctx) error {

	userRequest := new(request.UserUpdateRequest)

	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})

	}

	var user entity.User
	userID := ctx.Params("id")

	err := config.DB.First(&user, "id = ?", userID).Error

	if err != nil {

		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})

	}

	var errors []*request.ErrorResponse
	validate := validator.New()
	errValidate := validate.Struct(userRequest)

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

	// Update User

	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}

	user.Email = userRequest.Email
	user.PhoneNumber = userRequest.PhoneNumber

	errUpdate := config.DB.Save(&user).Error

	if errUpdate != nil {

		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserDelete(ctx *fiber.Ctx) error {

	var user entity.User
	userID := ctx.Params("id")

	err := config.DB.First(&user, "id = ?", userID).Error

	if err != nil {

		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})

	}

	errDelete := config.DB.Delete(&user).Error

	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Delete data success",
	})
}
