package database

import (
	"github.com/arshamalh/blogo/models"
)

func CheckUserExists(username string) bool {
	results := DB.Take(&models.User{
		Username: username,
	})
	return results.RowsAffected > 0
}

func CreateUser(username, password, email string) (uint, error) {
	user := models.User{
		Username: username,
		Password: password,
		Email:    email,
	}
	err := DB.Create(&user).Error
	return user.ID, err
}
