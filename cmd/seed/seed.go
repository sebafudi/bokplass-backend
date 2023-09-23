package main

import (
	"os"

	"github.com/sebafudi/bokplass-backend/internal/database"
	"github.com/sebafudi/bokplass-backend/internal/models"
)

func seedAuthors() {
	authors := []models.Author{
		{Name: "J.K. Rowling"},
		{Name: "J.R.R. Tolkien"},
		{Name: "George R.R. Martin"},
		{Name: "Stephen King"},
		{Name: "John E. Douglas"},
		{Name: "Mark Olshaker"},
	}
	for _, author := range authors {
		database.NewAuthor(&author)
	}
}

func seedBooks() {
	// book with Author ID
	books := []models.Book{
		{Title: "Harry Potter and the Philosopher's Stone", Authors: []models.Author{{ID: 1}}},
		{Title: "Harry Potter and the Chamber of Secrets", Authors: []models.Author{{ID: 1}}},
		{Title: "Harry Potter and the Prisoner of Azkaban", Authors: []models.Author{{ID: 1}}},
		{Title: "Harry Potter and the Goblet of Fire", Authors: []models.Author{{ID: 1}}},
		{Title: "Harry Potter and the Order of the Phoenix", Authors: []models.Author{{ID: 1}}},
		{Title: "Harry Potter and the Half-Blood Prince", Authors: []models.Author{{ID: 1}}},
		{Title: "Harry Potter and the Deathly Hallows", Authors: []models.Author{{ID: 1}}},
		{Title: "The Hobbit", Authors: []models.Author{{ID: 2}}},
		{Title: "The Lord of the Rings: The Fellowship of the Ring", Authors: []models.Author{{ID: 2}}},
		{Title: "The Lord of the Rings: The Two Towers", Authors: []models.Author{{ID: 2}}},
		{Title: "The Lord of the Rings: The Return of the King", Authors: []models.Author{{ID: 2}}},
		{Title: "A Game of Thrones", Authors: []models.Author{{ID: 3}}},
		{Title: "A Clash of Kings", Authors: []models.Author{{ID: 3}}},
		{Title: "A Storm of Swords", Authors: []models.Author{{ID: 3}}},
		{Title: "A Feast for Crows", Authors: []models.Author{{ID: 3}}},
		{Title: "A Dance with Dragons", Authors: []models.Author{{ID: 3}}},
		{Title: "The Shining", Authors: []models.Author{{ID: 4}}},
		{Title: "It", Authors: []models.Author{{ID: 4}}},
		{Title: "The Stand", Authors: []models.Author{{ID: 4}}},
		{Title: "Misery", Authors: []models.Author{{ID: 4}}},
		{Title: "The Dark Tower: The Gunslinger", Authors: []models.Author{{ID: 4}}},
		{Title: "The Dark Tower II: The Drawing of the Three", Authors: []models.Author{{ID: 4}}},
		{Title: "The Dark Tower III: The Waste Lands", Authors: []models.Author{{ID: 4}}},
		{Title: "The Dark Tower IV: Wizard and Glass", Authors: []models.Author{{ID: 4}}},
		{Title: "The Dark Tower V: Wolves of the Calla", Authors: []models.Author{{ID: 4}}},
		{Title: "The Dark Tower VI: Song of Susannah", Authors: []models.Author{{ID: 4}}},
		{Title: "The Dark Tower VII: The Dark Tower", Authors: []models.Author{{ID: 4}}},
		{Title: "Mindhunter: Inside the FBI's Elite Serial Crime Unit", Authors: []models.Author{{ID: 5}, {ID: 6}}},
	}
	for _, book := range books {
		database.NewBook(&book)
	}

}

func main() {
	os.Remove("test.db")
	database.Connect()
	seedAuthors()
	seedBooks()
}
