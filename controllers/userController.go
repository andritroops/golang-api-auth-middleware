package userController

import (
	"github.com/andritroops/go-latihan/config"
	"github.com/andritroops/go-latihan/models"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {

	var users []models.GetUser

	config.DB.Table("users").Find(&users)
	return c.Status(fiber.StatusOK).JSON(users)
}

func Store(c *fiber.Ctx) error {

	user := new(models.ValidateUser)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})

	}

	errors := models.ValidateStruct(*user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	config.DB.Table("users").Create(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
	})

}
