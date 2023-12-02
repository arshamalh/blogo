package database

import (
	"context"

	"github.com/arshamalh/blogo/models"
)

func (bdb *bundb) CreatePost(post *models.Post) (uint, error) {
	_, err := bdb.db.NewInsert().Model(post).Exec(context.Background())
	return post.ID, err
}

func (bdb *bundb) DeletePost(id uint) error {
	_, err := bdb.db.NewDelete().Model((*models.Post)(nil)).Where("id = ?", id).Exec(context.Background())
	return err
}

func (bdb *bundb) UpdatePost(post *models.Post) error {
	_, err := bdb.db.NewUpdate().Model(post).Exec(context.Background())
	return err
}

func (bdb *bundb) GetPost(id uint) (models.Post, error) {
	var post models.Post
	err := bdb.db.NewSelect().Model(&post).Relation("Author").Relation("Category").Relation("Comments").Where("id = ?", id).Scan(context.Background())
	return post, err
}

func (bdb *bundb) GetPosts() ([]models.Post, error) {
	var posts []models.Post
	err := bdb.db.NewSelect().Model(&posts).Relation("Author").Relation("Category").Scan(context.Background())
	return posts, err
}
