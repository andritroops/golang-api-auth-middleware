package userController

import (
	"log"

	"github.com/andritroops/go-latihan/config"
	"github.com/andritroops/go-latihan/models/entity"
	"github.com/andritroops/go-latihan/models/request"
	"github.com/andritroops/go-latihan/utils"
	"github.com/gofiber/fiber/v2"
)

func Index(ctx *fiber.Ctx) error {

	var users []entity.User

	result := config.DB.Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.Status(fiber.StatusOK).JSON(users)
}

func Store(ctx *fiber.Ctx) error {

	user := new(request.UserCreateRequest)

	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})

	}

	errors := request.ValidateStruct(*user)
	if errors != nil {
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
