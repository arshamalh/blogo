package routes

import (
	"github.com/arshamalh/blogo/controllers"
	"github.com/arshamalh/blogo/databases"
	"github.com/arshamalh/blogo/middlewares"
	"github.com/arshamalh/blogo/models/permissions"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine, db databases.Database) {
	router.Use(middlewares.CORSMiddleware())
	user_routes := router.Group("/api/v1/users")
	uc := controllers.NewUserController(db)
	{
		user_routes.POST("/register", uc.UserRegister)
		user_routes.POST("/login", uc.UserLogin)
		user_routes.POST("/logout", middlewares.RequireLogin, uc.UserLogout)
		user_routes.GET("/check_username", uc.CheckUsername)
		user_routes.GET("/id", middlewares.RequireLogin, uc.UserID)
		// Get & Update Profile and more
		// user_routes.GET("/info/:id", controllers.UserInfo)
	}

	post_routes := router.Group("api/v1/posts")
	pc := controllers.NewPostController(db)
	{
		post_routes.POST("/", middlewares.RequireLogin, middlewares.CheckPermissions(db, permissions.CreatePost), pc.CreatePost)
		post_routes.GET("/:id", pc.GetPost)
		post_routes.GET("/", pc.GetPosts)
		post_routes.PATCH("/:id", middlewares.RequireLogin, middlewares.CheckPermissions(db, permissions.EditPost), pc.UpdatePost)
		post_routes.DELETE("/:id", middlewares.RequireLogin, middlewares.CheckPermissions(db, permissions.DeletePost), pc.DeletePost)
	}

	comment_routes := router.Group("api/v1/comments")
	comment_controller := controllers.NewCommentController(db)
	{
		comment_routes.POST("/", middlewares.RequireLogin, comment_controller.CreateComment)
	}

	category_routes := router.Group("api/v1/categories")
	cc := controllers.NewCategoryController(db)
	{
		category_routes.POST("/", middlewares.RequireLogin, middlewares.CheckPermissions(db, permissions.CreateCategory), cc.CreateCategory)
		category_routes.GET("/:id", cc.GetCategory)
		category_routes.GET("/", cc.GetCategories)
		// category_routes.PATCH("/:id", middlewares.RequireLogin, middlewares.CheckPermissions(db, permissions.UpdateCategory), controllers.UpdateCategory)
		// category_routes.DELETE("/:id", middlewares.RequireLogin, middlewares.CheckPermissions(db, permissions.DeleteCategory), controllers.DeleteCategory)
	}

	role_routes := router.Group("api/v1/roles")
	rc := controllers.NewRoleController(db)
	{
		role_routes.POST("/", middlewares.RequireLogin, middlewares.CheckPermissions(db, permissions.CreateRole), rc.CreateRole)
		role_routes.PATCH("/", middlewares.RequireLogin, middlewares.CheckPermissions(db, permissions.UpdateRole), rc.UpdateRole)
		role_routes.GET("/:id", rc.GetRole)
		role_routes.GET("/", rc.GetRoles)
		role_routes.DELETE("/:id", middlewares.RequireLogin, middlewares.CheckPermissions(db, permissions.DeleteRole), rc.DeleteRole)
	}
}
