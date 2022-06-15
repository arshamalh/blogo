package database

import (
	"github.com/arshamalh/blogo/models"
	"github.com/arshamalh/blogo/models/permissions"
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

func GetUserPermissions(user_id uint) []permissions.Permission {
	var user models.User
	err := DB.Preload("Role").First(&user, user_id).Error
	if err != nil {
		return nil
	}
	return permissions.Decompress(user.Role.Permissions)

}
