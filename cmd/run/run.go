package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebafudi/bokplass-backend/internal/database"
	"github.com/sebafudi/bokplass-backend/routes"
)

func main() {
	database.Connect()

	app := fiber.New()
	routes.Setup(app)

	app.Listen(":3000")
}
