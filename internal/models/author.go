package models

import (
	"gorm.io/gorm"
)

// Author struct
type Author struct {
	gorm.Model
	ID    uint64 `gorm:"primaryKey"`
	Name  string
	Books []Book `gorm:"many2many:book_authors"`
}
