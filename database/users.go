package database

import (
	"github.com/arshamalh/blogo/models"
)

func CheckUserExists(username string) bool {
	results := DB.Take(&models.User{}, "username = ?", username)
	return results.RowsAffected > 0
}

func CreateUser(user *models.User) (uint, error) {
	err := DB.Create(&user).Error
	return user.ID, err
}

func GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	err := DB.First(&user, "username = ?", username).Error
	return user, err
}
