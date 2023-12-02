package cmd

import (
	"fmt"
	"os"

	"github.com/arshamalh/blogo/databases"
	"github.com/arshamalh/blogo/log"
	"github.com/arshamalh/blogo/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting the server...")

		// Load environment variables
		if err := godotenv.Load(); err != nil {
			log.Gl.Error(err.Error())
			return
		}

		// Database
		db, err := databases.ConnectDB()
		if err != nil {
			log.Gl.Error(err.Error())
			return
		}

		// Router
		router := echo.New()
		routes.InitializeRoutes(router, db, log.InitializeLogger())
		router.StaticFS("/", os.DirFS("./ui"))

		// Start the server
		if err := router.Start(":8080"); err != nil {
			log.Gl.Error(err.Error())
		}
	},
}
