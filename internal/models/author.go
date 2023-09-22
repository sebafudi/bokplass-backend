package models

import (
	"gorm.io/gorm"
)

// Author struct
type Author struct {
	gorm.Model
	Name  string
	Books []*Book `gorm:"many2many:book_authors"`
}
