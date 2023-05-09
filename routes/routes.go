package routes

import (
	"net/http"

	"github.com/arshamalh/blogo/controllers"
	"github.com/arshamalh/blogo/databases"
	"github.com/arshamalh/blogo/middlewares"
	"go.uber.org/zap"

	"github.com/arshamalh/blogo/models/permissions"
	"github.com/labstack/echo/v4"
	emdware "github.com/labstack/echo/v4/middleware"
)

func InitializeRoutes(router *echo.Echo, db databases.Database, logger *zap.Logger) {
	// Basic configurations and middleware.
	router.Use(emdware.CORSWithConfig(emdware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
	}))

	router.Use(middlewares.ZapLogger(logger))

	router.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusNoContent, nil)
	})

	/// ** Routes

	user_routes := router.Group("/api/v1/users")
	uc := controllers.NewUserController(db, logger)
	{
		user_routes.POST("/register", uc.UserRegister)
		user_routes.POST("/login", uc.UserLogin)
		user_routes.POST("/logout", uc.UserLogout, middlewares.RequireLogin)
		user_routes.GET("/check_username", uc.CheckUsername)
		user_routes.GET("/id", uc.UserID, middlewares.RequireLogin)
		// Get & Update Profile and more
		// user_routes.GET("/info/:id", controllers.UserInfo)
	}

	post_routes := router.Group("api/v1/posts")
	pc := controllers.NewPostController(db, logger)
	{
		post_routes.POST("/", pc.CreatePost, middlewares.RequireLogin, middlewares.CheckPermissions(db, permissions.CreatePost))
		post_routes.GET("/:id", pc.GetPost)
		post_routes.GET("/", pc.GetPosts)
		post_routes.PATCH("/:id", pc.UpdatePost, middlewares.RequireLogin, middlewares.CheckPermissions(db, permissions.EditPost))
		post_routes.DELETE("/:id", pc.DeletePost, middlewares.RequireLogin, middlewares.CheckPermissions(db, permissions.DeletePost))
	}

	comment_routes := router.Group("api/v1/comments")
	comment_controller := controllers.NewCommentController(db, logger)
	{
		comment_routes.POST("/", comment_controller.CreateComment, middlewares.RequireLogin)
	}

	category_routes := router.Group("api/v1/categories")
	cc := controllers.NewCategoryController(db, logger)
	{
		category_routes.POST("/", cc.CreateCategory, middlewares.RequireLogin, middlewares.CheckPermissions(db, permissions.CreateCategory))
		category_routes.GET("/:id", cc.GetCategory)
		category_routes.GET("/", cc.GetCategories)
		// category_routes.PATCH("/:id", middlewares.RequireLogin, middlewares.CheckPermissions(db, permissions.UpdateCategory), controllers.UpdateCategory)
		// category_routes.DELETE("/:id", middlewares.RequireLogin, middlewares.CheckPermissions(db, permissions.DeleteCategory), controllers.DeleteCategory)
	}

	role_routes := router.Group("api/v1/roles")
	rc := controllers.NewRoleController(db, logger)
	{
		role_routes.POST("/", rc.CreateRole, middlewares.RequireLogin, middlewares.CheckPermissions(db, permissions.CreateRole))
		role_routes.PATCH("/", rc.UpdateRole, middlewares.RequireLogin, middlewares.CheckPermissions(db, permissions.UpdateRole))
		role_routes.GET("/:id", rc.GetRole)
		role_routes.GET("/", rc.GetRoles)
		role_routes.DELETE("/:id", rc.DeleteRole, middlewares.RequireLogin, middlewares.CheckPermissions(db, permissions.DeleteRole))
	}
}
