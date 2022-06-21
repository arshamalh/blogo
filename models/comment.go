package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	User   User `json:"user"`
	UserID uint
	PostID uint   `json:"post_id" from:"post_id" binding:"required"`
	Text   string `gorm:"not null" json:"text" from:"text" binding:"required"`
}
