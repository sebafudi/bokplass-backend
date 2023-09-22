package database

import (
	"github.com/sebafudi/bokplass-backend/internal/models"
)

func GetBooks() []models.Book {
	var books []models.Book
	DB.Find(&books)
	return books
}

func GetBook(id uint64) (models.Book, error) {
	var book models.Book
	result := DB.First(&book, id)
	if result.Error != nil {
		return models.Book{}, result.Error
	}
	return book, nil
}

func NewBook(book *models.Book) {
	DB.Create(&book)
}

func UpdateBook(id uint64, book *models.Book) {
	DB.Model(&models.Book{}).Where("id = ?", id).Updates(book)
}

func DeleteBook(id uint64) {
	DB.Delete(&models.Book{}, id)
}
