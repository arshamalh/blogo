package main

import (
	"fmt"
	"log"
	"os"

	database "github.com/arshamalh/blogo/databases/gorm"
	"github.com/arshamalh/blogo/routes"
	"github.com/arshamalh/blogo/tools"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// @title Blogo API server
// @version 1.0
// @description A simple blog for educational purposes

// @host localhost:80
// @BasePath /api/v1
func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		fmt.Println("no environment variables", err)
	}

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

	if err := router.Start(":80"); err != nil {
		log.Fatal(err)
	}
}
