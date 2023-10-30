package databases

import (
	"github.com/arshamalh/blogo/models"
	"github.com/arshamalh/blogo/models/permissions"
)

type Database interface {
	// Users
	CheckUserExists(username string) bool
	CreateUser(user *models.User) (uint, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserPermissions(user_id uint) []permissions.Permission

	// Categories
	CheckCategoryExists(name string) bool
	CreateCategory(catg *models.Category) (uint, error)
	GetCategory(name string) (*models.Category, error)
	GetCategories() ([]models.Category, error)

	// Comments
	AddComment(comment *models.Comment) error
	GetComment(id uint) (*models.Comment, error)

	// Posts
	CreatePost(post *models.Post) (uint, error)
	DeletePost(id uint) error
	UpdatePost(post *models.Post) error
	GetPost(id uint) (models.Post, error)
	GetPosts() ([]models.Post, error)

	// Roles
	CreateRole(role *models.Role) error
	UpdateRole(role *models.Role) error
	DeleteRole(id uint) error
	GetRole(id uint) (models.Role, error)
	GetRoles() ([]models.Role, error)
}

/*

type Database struct {
	logger *log.Logger
}

func NewDatabase(logger *log.Logger) *Database {
	return &Database{logger: logger}
}

func (db *Database) CheckUserExists(username string) bool {
	db.logger.Println("CheckUserExists - Checking if user exists:", username)

	return true
}

func (db *Database) CreateUser(user *models.User) (uint, error) {
	db.logger.Println("CreateUser - Creating new user:", user.Username)

	return 1, nil
}

func (db *Database) GetUserByUsername(username string) (*models.User, error) {
	db.logger.Println("GetUserByUsername - Getting user by username:", username)

	return &models.User{}, nil
}

*/
