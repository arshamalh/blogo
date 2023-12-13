package models

type Category struct {
	ID          uint    `bun:"id"`
	Name        string  `json:"name" bun:"unique"`
	Description string  `json:"description"`
	Posts       []*Post `json:"posts" bun:"many2many:post_categories;join:post_category"`
}
