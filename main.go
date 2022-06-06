package main

// user routes and database
// register
// login
// logout

import (
	"blogo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.IntializeRoutes(router)
	router.Run(":8080")
}
