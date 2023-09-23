package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sebafudi/bokplass-backend/internal/database"
	"github.com/sebafudi/bokplass-backend/internal/models"
	"github.com/sebafudi/bokplass-backend/internal/views"
)

func GetBooks(c *fiber.Ctx) error {
	// search param
	search := c.Query("search")
	var books []models.Book
	if search != "" {
		books = database.SearchBooks(search)
	} else {
		books = database.GetBooks()
	}
	for i, book := range books {
		books[i].Authors = database.GetAuthorsByBook(book.ID)
	}
	return c.Type("html").SendString(views.RenderBooks(books))
}

func GetBook(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}
	var book models.Book
	book, err = database.GetBook(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Book not found",
		})
	}
	return c.JSON(book)
}

func NewBook(c *fiber.Ctx) error {
	var book models.Book
	err := c.BodyParser(&book)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	database.NewBook(&book)
	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}
	var book models.Book
	err = c.BodyParser(&book)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	database.UpdateBook(id, &book)
	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}
	database.DeleteBook(id)
	return nil
}
