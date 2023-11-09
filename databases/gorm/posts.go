package database

import (
	"fmt"

	"github.com/arshamalh/blogo/log"
	"github.com/arshamalh/blogo/models"
)

func (gdb *gormdb) CreatePost(post *models.Post) (uint, error) {
	err := gdb.db.Create(post).Error

	if err == nil {
		log.Gl.Info(fmt.Sprintf("Post created with ID: %d", post.ID))
	}
	return post.ID, err
}

func (gdb *gormdb) DeletePost(id uint) error {
	err := gdb.db.Delete(&models.Post{}, id).Error

	if err == nil {
		log.Gl.Info(fmt.Sprintf("Post with ID %d deleted", id))
	}

	return err
}

func (gdb *gormdb) UpdatePost(post *models.Post) error {
	err := gdb.db.Save(post).Error

	if err == nil {
		log.Gl.Info(fmt.Sprintf("Post with ID %d updated", post.ID))
	}
	return err
}

func (gdb *gormdb) GetPost(id uint) (models.Post, error) {
	var post models.Post
	err := gdb.db.Preload("Author").Preload("Category").Preload("Comment").First(&post, id).Error

	if err != nil {
		log.Gl.Error(fmt.Sprintf("Error: %d: %v", id, err))
	}

	return post, err
}

func (gdb *gormdb) GetPosts() ([]models.Post, error) {
	var posts []models.Post
	err := gdb.db.Preload("Author").Preload("Category").Find(&posts).Error

	if err != nil {
		log.Gl.Error(fmt.Sprintf("Error: %v", err))
	}

	return posts, err
}
