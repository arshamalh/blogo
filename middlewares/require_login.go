package middlewares

import (
	"net/http"
	"os"

	"github.com/arshamalh/blogo/log"
	"github.com/arshamalh/blogo/tools"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func RequireLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		access_token, err := ctx.Cookie("access_token")
		if err != nil || access_token.Value == "" {
			log.Gl.Error("Error:", zap.Any("error", err))
			return ctx.JSON(http.StatusUnauthorized, "you should log in")
		}

		jwt_access, err := tools.ExtractTokenData(access_token.Value, os.Getenv("JWT_SECRET"))
		if err != nil {
			log.Gl.Error("Error:", zap.Any("error", err))
			return ctx.JSON(http.StatusUnauthorized, "invalid token")
		}

		// If access token is valid and not expired, extract data from it
		userID := jwt_access.Subject
		ctx.Set("user_id", userID)
		log.Gl.Info("User with ID has been authenticated", zap.Any("user_id", userID))
		return next(ctx)
	}
}
