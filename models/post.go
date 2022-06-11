package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	AuthorID   uint       `json:"author_id"`
	Author     User       `json:"author"`
	Categories []Category `json:"categories" gorm:"many2many:post_categories;"`
}
