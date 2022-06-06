package routes

import (
	"github.com/arshamalh/blogo/controllers"
	"github.com/gin-gonic/gin"
)

func IntializeRoutes(router *gin.Engine) error {
	user_routes := router.Group("/api/v1/users")
	{
		user_routes.POST("/register", controllers.UserRegister)
		user_routes.POST("/login", controllers.UserLogin)
		user_routes.POST("/logout", controllers.UserLogout)
		user_routes.GET("/check_username", controllers.CheckUsername)
		// Get & Update Profile and more
	}
	return nil
}
