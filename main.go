// main.go
package main

import (
	"os"

	"github.com/arshamalh/blogo/log"
	"github.com/arshamalh/blogo/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Blogo API server
// @version 1.0
// @description A simple blog for educational purposes
// @contact.email hadirezaei.etc.esfahan@gmail.com
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Load environment variables
	logger := log.InitializeLogger()

	if err := godotenv.Load(); err != nil {
		log.Gl.Error(err.Error())
	}

	// Database
	db, err := databases.Connect()
	if err != nil {
		log.Gl.Error(err.Error())
		return
	}

	// Router
	router := echo.New()
	routes.InitializeRoutes(router, db, logger)

	// Swagger
	router.GET("/swagger/*", echoSwagger.WrapHandler)

	router.StaticFS("/", os.DirFS("./ui"))
	if err := router.Start(":8080"); err != nil {
		log.Gl.Error(err.Error())
	}

}
