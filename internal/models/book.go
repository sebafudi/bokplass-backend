package models

import (
	"gorm.io/gorm"
)

// Book struct
type Book struct {
	gorm.Model
	Title   string
	Authors []*Author `gorm:"many2many:book_authors"`
}
