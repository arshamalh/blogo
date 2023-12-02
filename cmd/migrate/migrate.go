package cmd

import (
	"fmt"

	"github.com/arshamalh/blogo/databases"
	"github.com/arshamalh/blogo/models"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Migrating the database...")
		db, err := databases.ConnectDB()
		if err != nil {
			fmt.Println("Error connecting to the database:", err)
			return
		}

		// Check if the database supports migrations
		migrator, ok := db.(databases.Migrator)
		if !ok {
			fmt.Println("Database does not support migrations")
			return
		}

		// Bun Auto Migration
		err = migrator.AutoMigrate(
			&models.User{},
			&models.Post{},
			&models.Category{},
			&models.Role{},
			&models.Comment{},
		)
		if err != nil {
			fmt.Println("Error performing migration:", err)
			return
		}

		fmt.Println("Database migration completed.")
	},
}
