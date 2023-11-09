package routes

import (
	"net/http"

	"github.com/arshamalh/blogo/controllers"
	"github.com/arshamalh/blogo/databases"
	"github.com/arshamalh/blogo/log"
	"go.uber.org/zap"

	"github.com/labstack/echo/v4"
	emdware "github.com/labstack/echo/v4/middleware"
)

func InitializeRoutes(router *echo.Echo, db databases.Database, logger *zap.Logger) {

	// Basic configurations and middleware.
	router.Use(emdware.CORSWithConfig(emdware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
	}))

	router.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusNoContent, nil)
	})

	// Routes for users
	userRoutes := router.Group("/api/v1/users")
	uc := controllers.NewUserController(db, logger)
	{
		userRoutes.POST("/register", func(c echo.Context) error {
			log.Gl.Info("User registration request received")
			return uc.UserRegister(c)
		})

	}

	// Routes for posts
	postRoutes := router.Group("/api/v1/posts")
	pc := controllers.NewPostController(db, logger)
	{
		postRoutes.POST("/", func(c echo.Context) error {
			log.Gl.Info("Post creation request received")
			return pc.CreatePost(c)
		})

		postRoutes.PATCH("/:id", func(c echo.Context) error {
			log.Gl.Info("Post update request received")
			return pc.UpdatePost(c)
		})
	}

	// Routes for roles
	roleRoutes := router.Group("/api/v1/roles")
	rc := controllers.NewRoleController(db, logger)
	{
		roleRoutes.POST("/", func(c echo.Context) error {
			log.Gl.Info("Role creation request received")
			return rc.CreateRole(c)
		})

		roleRoutes.DELETE("/:id", func(c echo.Context) error {
			log.Gl.Info("Role deletion request received")
			return rc.DeleteRole(c)
		})
	}
}
