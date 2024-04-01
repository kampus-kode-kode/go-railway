package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kampus-kode-kode/go-railway/routes"
)

func main() {
	app := fiber.New()

	routes.SetupTaskRoutes(app)

	app.Listen(":3000")
}
