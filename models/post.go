package models

import (
	"github.com/uptrace/bun"
)

type Post struct {
	ID            uint `bun:"id"`
	bun.BaseModel `bun:"post"`
	Title         string      `json:"title"`
	Content       string      `json:"content"`
	AuthorID      uint        `json:"author_id"`
	Author        *User       `json:"author" bun:"rel:belongs-to"`
	Comments      []*Comment  `json:"comments" bun:"rel:has-many"`
	Categories    []*Category `json:"categories" bun:"many2many:post_categories;join:post_category"`
}
