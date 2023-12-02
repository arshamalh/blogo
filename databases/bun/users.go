package database

import (
	"context"

	"github.com/arshamalh/blogo/models"
	"github.com/arshamalh/blogo/models/permissions"
)

func (bdb *bundb) CheckUserExists(username string) bool {
	var count int
	_, err := bdb.db.NewSelect().Model((*models.User)(nil)).Where("username = ?", username).Count(context.Background())
	if err != nil {
		return false
	}
	return count > 0
}

func (bdb *bundb) CreateUser(user *models.User) (uint, error) {
	_, err := bdb.db.NewInsert().Model(user).Exec(context.Background())
	return user.ID, err
}

func (bdb *bundb) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := bdb.db.NewSelect().Model(&user).Where("username = ?", username).Scan(context.Background())
	return &user, err
}

func (bdb *bundb) GetUserPermissions(user_id uint) []permissions.Permission {
	var user models.User
	err := bdb.db.NewSelect().Model(&user).Relation("Role").Where("id = ?", user_id).Scan(context.Background())
	if err != nil {
		return nil
	}
	return permissions.Decompress(user.Role.Permissions)
}
