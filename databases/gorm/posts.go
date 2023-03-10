package database

import "github.com/arshamalh/blogo/models"

func (gdb *gormdb) CreatePost(post *models.Post) (uint, error) {
	err := gdb.db.Create(&post).Error
	return post.ID, err
}

func (gdb *gormdb) DeletePost(id uint) error {
	return gdb.db.Delete(&models.Post{}, "id = ?", id).Error
}

func (gdb *gormdb) UpdatePost(post *models.Post) error {
	return gdb.db.Updates(&post).Error
}

func (gdb *gormdb) GetPost(id uint) (models.Post, error) {
	// FIXME: This is not the best way to do this
	var post models.Post
	err := gdb.db.
		Preload("Author").
		Preload("Category").
		Preload("Comment").
		First(&post, id).Error
	return post, err
}

func (gdb *gormdb) GetPosts() ([]models.Post, error) {
	var posts []models.Post
	err := gdb.db.Preload("Author").Preload("Category").Find(&posts).Error
	return posts, err
}
