package database

import (
	"fmt"

	"github.com/arshamalh/blogo/log"
	"github.com/arshamalh/blogo/models"
	"github.com/arshamalh/blogo/models/permissions"
)

func (gdb *gormdb) CheckUserExists(username string) bool {
	results := gdb.db.Take(&models.User{}, "username = ?", username)
	return results.RowsAffected > 0
}

func (gdb *gormdb) CreateUser(user *models.User) (uint, error) {
	err := gdb.db.Create(&user).Error

	if err == nil {
		log.Gl.Error(fmt.Sprintf("Failed to get permissions for user with ID: %d Error: %s", user.ID, err))
	}

	return user.ID, err
}

func (gdb *gormdb) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	err := gdb.db.First(&user, "username = ?", username).Error

	if err != nil {
		log.Gl.Error(fmt.Sprintf("Failed to get user with username %s Error: %s", username, err))

	}

	return user, err
}

func (gdb *gormdb) GetUserPermissions(user_id uint) []permissions.Permission {
	var user models.User
	err := gdb.db.Preload("Role").First(&user, user_id).Error

	if err != nil {
		log.Gl.Error(fmt.Sprintf("Failed to get permissions for user with ID: %d Error: %s", user_id, err))

		return nil
	}

	return permissions.Decompress(user.Role.Permissions)
}
