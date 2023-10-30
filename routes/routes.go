package routes

import (
	"log"
	"net/http"

	"github.com/arshamalh/blogo/controllers"
	"github.com/arshamalh/blogo/databases"
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
		log.Printf("Health check request received")
		return c.JSON(http.StatusNoContent, nil)
	})

	/// ** Routes

	user_routes := router.Group("/api/v1/users")
	uc := controllers.NewUserController(db, logger)
	{
		user_routes.POST("/register", func(c echo.Context) error {
			log.Printf("User registration request received")
			return uc.UserRegister(c)
		})

		user_routes.POST("/logout", func(c echo.Context) error {
			log.Printf("User logout request received")
			return uc.UserLogout(c)
		})
	}

	post_routes := router.Group("api/v1/posts")
	pc := controllers.NewPostController(db, logger)
	{
		post_routes.POST("/", func(c echo.Context) error {
			log.Printf("Post creation request received")
			return pc.CreatePost(c)
		})

		post_routes.PATCH("/:id", func(c echo.Context) error {
			log.Printf("Post update request received")
			return pc.UpdatePost(c)
		})

		role_routes := router.Group("api/v1/roles")
		rc := controllers.NewRoleController(db, logger)
		{
			role_routes.POST("/", func(c echo.Context) error {
				log.Printf("Role creation request received")
				return rc.CreateRole(c)
			})

			role_routes.DELETE("/:id", func(c echo.Context) error {
				log.Printf("Role deletion request received")
				return rc.DeleteRole(c)
			})
		}
	}
}
