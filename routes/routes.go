package routes

import (
	"net/http"

	"github.com/arshamalh/blogo/controllers"
	"github.com/arshamalh/blogo/databases"
	"github.com/arshamalh/blogo/middlewares"
	"go.uber.org/zap"

	_ "github.com/arshamalh/blogo/docs"
	"github.com/arshamalh/blogo/models/permissions"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitializeRoutes(router *echo.Echo, db databases.Database, logger *zap.Logger) {
	// Basic configurations and middleware.
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	router.Use(middlewares.ZapLogger(logger))

	router.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusNoContent, nil)
	})

	router.GET("/docs/*", echoSwagger.WrapHandler)

	/// ** Routes

	userRoutes := router.Group("/api/v1/users")
	uc := controllers.NewUserController(db, logger)
	{
		userRoutes.POST("/register", uc.UserRegister)
		userRoutes.POST("/login", uc.UserLogin)
		userRoutes.POST("/logout", uc.UserLogout, middlewares.RequireLogin)
		userRoutes.GET("/check_username", uc.CheckUsername)
		userRoutes.GET("/id", uc.UserID, middlewares.RequireLogin)
		// Get & Update Profile and more
		// userRoutes.GET("/info/:id", controllers.UserInfo)
	}

	postRoutes := router.Group("api/v1/posts")
	pc := controllers.NewPostController(db, logger)
	{
		postRoutes.POST("/", pc.CreatePost, middlewares.RequireLogin, middlewares.CheckPermissions(db, permissions.CreatePost))
		postRoutes.GET("/:id", pc.GetPost)
		postRoutes.GET("/", pc.GetPosts)
		postRoutes.PATCH("/:id", pc.UpdatePost, middlewares.RequireLogin, middlewares.CheckPermissions(db, permissions.EditPost))
		postRoutes.DELETE("/:id", pc.DeletePost, middlewares.RequireLogin, middlewares.CheckPermissions(db, permissions.DeletePost))
	}

	commentRoutes := router.Group("api/v1/comments")
	commentController := controllers.NewCommentController(db)
	{
		commentRoutes.POST("/", commentController.CreateComment, middlewares.RequireLogin)
	}

	categoryRoutes := router.Group("api/v1/categories")
	cc := controllers.NewCategoryController(db, logger)
	{
		categoryRoutes.POST("/", cc.CreateCategory, middlewares.RequireLogin, middlewares.CheckPermissions(db, permissions.CreateCategory))
		categoryRoutes.GET("/:id", cc.GetCategory)
		categoryRoutes.GET("/", cc.GetCategories)
		// categoryRoutes.PATCH("/:id", middlewares.RequireLogin, middlewares.CheckPermissions(db, permissions.UpdateCategory), controllers.UpdateCategory)
		// categoryRoutes.DELETE("/:id", middlewares.RequireLogin, middlewares.CheckPermissions(db, permissions.DeleteCategory), controllers.DeleteCategory)
	}

	roleRoutes := router.Group("api/v1/roles")
	rc := controllers.NewRoleController(db, logger)
	{
		roleRoutes.POST("/", rc.CreateRole, middlewares.RequireLogin, middlewares.CheckPermissions(db, permissions.CreateRole))
		roleRoutes.PATCH("/", rc.UpdateRole, middlewares.RequireLogin, middlewares.CheckPermissions(db, permissions.UpdateRole))
		roleRoutes.GET("/:id", rc.GetRole)
		roleRoutes.GET("/", rc.GetRoles)
		roleRoutes.DELETE("/:id", rc.DeleteRole, middlewares.RequireLogin, middlewares.CheckPermissions(db, permissions.DeleteRole))
	}
}
