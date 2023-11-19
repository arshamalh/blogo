package database

import (
	"context"

	"github.com/arshamalh/blogo/models"
)

func (bdb *bundb) AddComment(comment *models.Comment) error {
	_, err := bdb.db.NewInsert().Model(comment).Exec(context.Background())
	return err
}

func (bdb *bundb) GetComment(id uint) (*models.Comment, error) {
	var comment models.Comment
	err := bdb.db.NewSelect().Model(&comment).Where("id = ?", id).Scan(context.Background())
	return &comment, err
}
