package database

import (
	"fmt"

	"github.com/arshamalh/blogo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dsn string) {
	DBConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")

	// Auto migration
	err = DBConnection.AutoMigrate(models.User{})
	if err != nil {
		panic("Failed to migrate the database")
	}
	fmt.Println("Database Migrated")
}
