package database

import (
	"fmt"

	"github.com/arshamalh/blogo/models"
	"github.com/arshamalh/blogo/models/permissions"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dsn string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")

	// Auto migration
	err = DB.AutoMigrate(models.User{}, models.Post{}, models.Category{}, models.Role{}, models.Comment{})
	if err != nil {
		panic("Failed to migrate the database")
	}
	AddBasicRoles()
	fmt.Println("Database Migrated")
}

// Add some basic roles manually
func AddBasicRoles() {
	CreateRole(&models.Role{Name: "superadmin", Permissions: permissions.Compress([]permissions.Permission{permissions.FullAccess})})
	CreateRole(&models.Role{Name: "moderator", Permissions: permissions.Compress([]permissions.Permission{permissions.FullContents})})
	CreateRole(&models.Role{Name: "author", Permissions: permissions.Compress([]permissions.Permission{permissions.CreatePost, permissions.FullContents})})
}
