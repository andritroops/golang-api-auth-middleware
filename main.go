package main

import (
	"github.com/andritroops/go-latihan/config"
	userController "github.com/andritroops/go-latihan/controllers"
	"github.com/gofiber/fiber/v2"
)

func main() {

	config.ConnectDatabase()

	app := fiber.New()

	app.Get("/api/users", userController.Index)
	app.Post("/api/users", userController.Store)

	app.Listen("127.0.0.1:3000")

}
