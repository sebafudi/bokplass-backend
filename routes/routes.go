package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebafudi/bokplass-backend/internal/controllers"
)

func Setup(app *fiber.App) {
	// static
	app.Static("/", "./public")

	app.Get("/api/v1/authors", controllers.GetAuthors)
	app.Get("/api/v1/authors/:id", controllers.GetAuthor)
	app.Post("/api/v1/authors", controllers.NewAuthor)
	app.Put("/api/v1/authors/:id", controllers.UpdateAuthor)
	app.Delete("/api/v1/authors/:id", controllers.DeleteAuthor)

	app.Get("/api/v1/books", controllers.GetBooks)
	app.Get("/api/v1/books/:id", controllers.GetBook)
	app.Post("/api/v1/books", controllers.NewBook)
	app.Put("/api/v1/books/:id", controllers.UpdateBook)
	app.Delete("/api/v1/books/:id", controllers.DeleteBook)
}
