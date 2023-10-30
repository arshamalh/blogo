package database

import "github.com/arshamalh/blogo/models"

func (gdb *gormdb) AddComment(comment *models.Comment) error {
	return gdb.db.Create(comment).Error
}

func (gdb *gormdb) GetComment(id uint) (*models.Comment, error) {
	var comment models.Comment
	err := gdb.db.Where("id = ?", id).First(&comment).Error
	return &comment, err
}
