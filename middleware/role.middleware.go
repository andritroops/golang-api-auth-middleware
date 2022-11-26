package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Role(ctx *fiber.Ctx) error {

	is_admin := ctx.Locals("role")

	if is_admin == "1" {
		return ctx.Next()
	}

	log.Println(is_admin)
	return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"message": "Access Forbidden",
	})

}
