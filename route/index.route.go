package route

import (
	"github.com/andritroops/go-latihan/controllers"
	"github.com/andritroops/go-latihan/middleware"
	"github.com/andritroops/go-latihan/utils"
	"github.com/gofiber/fiber/v2"
)

func Route(route *fiber.App) {
	route.Post("/api/auth/signin", controllers.Signin)
	route.Get("/api/users", middleware.Auth, middleware.Role, controllers.UserIndex)
	route.Post("/api/users", middleware.Auth, middleware.Role, utils.HandleSingleFIle, controllers.UserStore)
	route.Get("/api/users/:id", middleware.Auth, middleware.Role, controllers.UserEdit)
	route.Put("/api/users/:id", middleware.Auth, middleware.Role, controllers.UserUpdate)
	route.Delete("/api/users/:id", middleware.Auth, middleware.Role, controllers.UserDelete)

	route.Post("/api/categories", middleware.Auth, middleware.Role, utils.HandleMultipleFIle, controllers.CategoryStore)
}
