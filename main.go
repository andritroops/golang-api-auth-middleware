package main

import (
	"github.com/andritroops/go-latihan/config"
	"github.com/andritroops/go-latihan/route"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initial Database
	config.ConnectDatabase()

	// Run Migration
	config.RunMigration()

	app := fiber.New()

	// Init Route
	route.Route(app)

	app.Listen("127.0.0.1:3000")

}
