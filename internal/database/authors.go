package database

import (
	"github.com/sebafudi/bokplass-backend/internal/models"
)

func GetAuthors() []models.Author {
	var authors []models.Author
	DB.Find(&authors)
	return authors
}

func GetAuthor(id string) models.Author {
	var author models.Author
	DB.Find(&author, id)
	return author
}

func NewAuthor(author *models.Author) {
	DB.Create(&author)
}

func UpdateAuthor(id string, author *models.Author) {
	DB.Model(&models.Author{}).Where("id = ?", id).Updates(author)
}

func DeleteAuthor(id string) {
	DB.Delete(&models.Author{}, id)
}
