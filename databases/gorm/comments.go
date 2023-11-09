package database

import (
	"github.com/arshamalh/blogo/log"
	"github.com/arshamalh/blogo/models"
	"go.uber.org/zap"
)

func (gdb *gormdb) AddComment(comment *models.Comment) error {
	err := gdb.db.Create(comment).Error

	if err == nil {
		log.Gl.Info("Comment added")
	}

	return err
}

func (gdb *gormdb) GetComment(id uint) (*models.Comment, error) {
	var comment models.Comment
	err := gdb.db.Where("id = ?", id).First(&comment).Error

	if err == nil {
		log.Gl.Info("Retrieved comment with ID", zap.Uint("ID", id))
	}

	return &comment, err
}
