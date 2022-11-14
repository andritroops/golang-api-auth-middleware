package route

import (
	userController "github.com/andritroops/go-latihan/controllers"
	"github.com/andritroops/go-latihan/middleware"
	"github.com/gofiber/fiber/v2"
)

func Route(route *fiber.App) {
	route.Get("/api/users", middleware.Auth, userController.Index)
	route.Post("/api/users", userController.Store)
}
