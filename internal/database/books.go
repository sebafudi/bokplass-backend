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

func GetBooksByAuthor(authorID uint64) []models.Book {
	var books []models.Book
	DB.Joins("JOIN book_authors ON books.id = book_authors.book_id").
		Where("book_authors.author_id = ?", authorID).
		Find(&books)
	return books
}

func SearchBooks(search string) []models.Book {
	var books []models.Book
	DB.Where("LOWER(title) LIKE LOWER(?)", "%"+search+"%").Find(&books)
	return books
}
