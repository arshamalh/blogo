package models

import "github.com/uptrace/bun"

type Category struct {
	ID            uint `bun:"id"`
	bun.BaseModel `bun:"category"`
	Name          string  `json:"name" bun:"unique"`
	Description   string  `json:"description"`
	Posts         []*Post `json:"posts" bun:"many2many:post_categories;join:post_category"`
}
