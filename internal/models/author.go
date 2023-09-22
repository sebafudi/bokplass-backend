package models

import (
	"gorm.io/gorm"
)

// Author struct
type Author struct {
	gorm.Model
	ID    uint `gorm:"primaryKey"`
	Name  string
	Books []Book `gorm:"many2many:book_authors"`
}
