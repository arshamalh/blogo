package main

import (
	"os"

	database "github.com/arshamalh/blogo/databases/gorm"
	"github.com/arshamalh/blogo/log"
	"github.com/arshamalh/blogo/routes"
	"github.com/arshamalh/blogo/tools"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load environment variables
	logger := log.InitializeLogger()

	if err := godotenv.Load(); err != nil {
		logger.Error(err.Error())
	}

	// Database
	dsn := tools.DBConfig{
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		Host:     os.Getenv("HOST"),
	}
	db, err := database.Connect(dsn.String())
	if err != nil {
		log.Gl.Error(err.Error())
		return
	}

	// Router
	router := echo.New()
	routes.InitializeRoutes(router, db, logger)
	router.StaticFS("/", os.DirFS("./ui"))

	router.Start(":80")
}

// @title Blogo API server
// @version 1.0
// @description A simple blog for educational purposes
// @host localhost:80
// @BasePath /api/v1
