package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `json:"name" gorm:"uniqueIndex"`
	Description string `json:"description"`
	Posts       []Post `json:"posts" gorm:"many2many:post_categories;"`
}
