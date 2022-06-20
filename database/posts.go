package database

import "github.com/arshamalh/blogo/models"

func CreatePost(post *models.Post) (uint, error) {
	err := DB.Create(&post).Error
	return post.ID, err
}

func DeletePost(id uint) error {
	return DB.Delete(&models.Post{}, "id = ?", id).Error
}

func UpdatePost(post *models.Post) error {
	return DB.Updates(&post).Error
}

func GetPost(id uint) (models.Post, error) {
	// FIXME: This is not the best way to do this
	var post models.Post
	err := DB.
		Preload("Author").
		Preload("Category").
		Preload("Comment").
		First(&post, id).Error
	return post, err
}

func GetPosts() ([]models.Post, error) {
	var posts []models.Post
	err := DB.Preload("Author").Preload("Category").Find(&posts).Error
	return posts, err
}
