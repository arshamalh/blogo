package database

import (
	"fmt"

	"github.com/arshamalh/blogo/models"
	"github.com/arshamalh/blogo/models/permissions"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type gormdb struct {
	db *gorm.DB
}

func Connect(dsn string) *gormdb {
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
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
	gdb := &gormdb{db: DB}
	gdb.AddBasicRoles()
	fmt.Println("Database Migrated")
	return gdb
}

// Add some basic roles manually
func (gdb *gormdb) AddBasicRoles() {
	gdb.CreateRole(&models.Role{Name: "superadmin", Permissions: permissions.Compress([]permissions.Permission{permissions.FullAccess})})
	gdb.CreateRole(&models.Role{Name: "moderator", Permissions: permissions.Compress([]permissions.Permission{permissions.FullContents})})
	gdb.CreateRole(&models.Role{Name: "author", Permissions: permissions.Compress([]permissions.Permission{permissions.CreatePost, permissions.FullContents})})
}
