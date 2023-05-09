package main

import (
	"os"

	database "github.com/arshamalh/blogo/databases/gorm"
	"github.com/arshamalh/blogo/routes"
	"github.com/arshamalh/blogo/tools"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load environment variables
	godotenv.Load()

	logger := tools.InitializeLogger()

	// Database
	dsn := tools.DBConfig{
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		Host:     os.Getenv("HOST"),
	}
	db := database.Connect(dsn.String())

	// Router
	router := echo.New()
	routes.InitializeRoutes(router, db, logger)
	router.StaticFS("/", os.DirFS("./ui"))

	router.Start(":80")
}
