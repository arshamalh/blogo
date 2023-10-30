package database

import (
	"log"

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
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connection successfully opened")

	// Auto migration
	err = DB.AutoMigrate(models.User{}, models.Post{}, models.Category{}, models.Role{}, models.Comment{})
	if err != nil {
		log.Fatalf("Failed to migrate the database: %v", err)
	}
	log.Println("Database Migrated")
	gdb := &gormdb{db: DB}
	gdb.AddBasicRoles()
	return gdb
}

// Add some basic roles manually
func (gdb *gormdb) AddBasicRoles() {
	if err := gdb.CreateRole(&models.Role{Name: "superadmin", Permissions: permissions.Compress([]permissions.Permission{permissions.FullAccess})}); err != nil {
		log.Printf("Failed to create role 'superadmin': %v", err)
	}
	if err := gdb.CreateRole(&models.Role{Name: "moderator", Permissions: permissions.Compress([]permissions.Permission{permissions.FullContents})}); err != nil {
		log.Printf("Failed to create role 'moderator': %v", err)
	}
	if err := gdb.CreateRole(&models.Role{Name: "author", Permissions: permissions.Compress([]permissions.Permission{permissions.CreatePost, permissions.FullContents})}); err != nil {
		log.Printf("Failed to create role 'author': %v", err)
	}
}

// CreateRole creates a role
func (gdb *gormdb) CreateRole(role *models.Role) error {
	err := gdb.db.Create(role).Error
	if err == nil {
		log.Printf("Role '%s' created successfully", role.Name)
	} else {
		log.Printf("Failed to create role '%s': %v", role.Name, err)
	}
	return err
}
