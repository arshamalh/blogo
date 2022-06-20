package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	User   User `json:"user"`
	UserID uint
	PostID uint
	Text   string `gorm:"not null"`
}
