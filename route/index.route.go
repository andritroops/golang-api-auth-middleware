package route

import (
	"github.com/andritroops/go-latihan/controllers"
	"github.com/andritroops/go-latihan/middleware"
	"github.com/gofiber/fiber/v2"
)

func Route(route *fiber.App) {
	route.Post("/api/auth/signin", controllers.Signin)
	route.Get("/api/users", middleware.Auth, controllers.UserIndex)
	route.Post("/api/users", middleware.Auth, controllers.UserStore)
	route.Get("/api/users/:id", middleware.Auth, controllers.UserEdit)
	route.Put("/api/users/:id", controllers.UserUpdate)
	route.Delete("/api/users/:id", controllers.UserDelete)
}
