package routes

import (
	"github.com/arshamalh/blogo/controllers"
	"github.com/arshamalh/blogo/middlewares"
	"github.com/gin-gonic/gin"
)

func IntializeRoutes(router *gin.Engine) {
	user_routes := router.Group("/api/v1/users")
	{
		user_routes.POST("/register", controllers.UserRegister)
		user_routes.POST("/login", controllers.UserLogin)
		user_routes.POST("/logout", controllers.UserLogout)
		user_routes.GET("/check_username", controllers.CheckUsername)
		// Get & Update Profile and more
		// Show Post creation page only to logged in users,
		//   so there should be an endpoint just to check logged in status
	}

	post_routes := router.Group("api/v1/posts")
	{
		post_routes.POST("/", middlewares.IsLoggedIn, controllers.CreatePost)
		// post_routes.GET("/:id", controllers.GetPost)
		// post_routes.GET("/", controllers.GetPosts)
	}
}
