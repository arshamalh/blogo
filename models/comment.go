package models

import (
	"github.com/uptrace/bun"
)

type Comment struct {
	bun.BaseModel `bun:"comment"`
	User          *User  `json:"user" bun:"rel:belongs-to"`
	UserID        uint   `json:"-" bun:"notnull"`
	PostID        uint   `json:"post_id" bun:"notnull"`
	Text          string `json:"text" bun:"notnull"`
}
