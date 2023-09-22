package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebafudi/bokplass-backend/internal/database"
	"github.com/sebafudi/bokplass-backend/internal/models"
)

func GetAuthors(c *fiber.Ctx) error {
	return c.JSON(database.GetAuthors())
}

func GetAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(database.GetAuthor(id))
}

func NewAuthor(c *fiber.Ctx) error {
	var author models.Author
	err := c.BodyParser(&author)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	database.NewAuthor(&author)
	return c.JSON(author)
}

func UpdateAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	var author models.Author
	err := c.BodyParser(&author)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	database.UpdateAuthor(id, &author)
	return c.JSON(author)
}

func DeleteAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DeleteAuthor(id)
	return nil
}
