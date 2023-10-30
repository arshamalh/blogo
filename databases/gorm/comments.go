package database

import (
	"log"

	"github.com/arshamalh/blogo/models"
)

func (gdb *gormdb) AddComment(comment *models.Comment) error {
	err := gdb.db.Create(comment).Error

	if err == nil {
		log.Printf("Added comment with ID: %d", comment.ID)
	} else {
		log.Printf("Failed to add comment. Error: %v", err)
	}

	return err
}

func (gdb *gormdb) GetComment(id uint) (*models.Comment, error) {
	var comment models.Comment
	err := gdb.db.Where("id = ?", id).First(&comment).Error

	if err == nil {
		log.Printf("Retrieved comment with ID: %d", id)
	} else {
		log.Printf("Failed to retrieve comment with ID: %d. Error: %v", id, err)
	}

	return &comment, err
}
