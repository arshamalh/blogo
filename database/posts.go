package database

import "github.com/arshamalh/blogo/models"

func CreatePost(post *models.Post) (uint, error) {
	err := DB.Create(&post).Error
	return post.ID, err
}

func GetPost(id uint) (models.Post, error) {
	var post models.Post
	err := DB.Preload("Author").Preload("Categories").First(&post, id).Error
	return post, err
}

func GetPosts() ([]models.Post, error) {
	var posts []models.Post
	err := DB.Preload("Author").Preload("Categories").Find(&posts).Error
	return posts, err
}
