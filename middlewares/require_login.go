package middlewares

import (
	"net/http"
	"os"

	"github.com/arshamalh/blogo/tools"
	"github.com/labstack/echo/v4"
)

func RequireLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		access_token, err := ctx.Cookie("access_token")
		if err != nil || access_token.Value == "" {
			return ctx.JSON(http.StatusUnauthorized, "you should login")
		}

		jwt_access, err := tools.ExtractTokenData(access_token.Value, os.Getenv("JWT_SECRET"))
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, "invalid token")
		}

		// If access token is valid and not expired, extract data from it
		userID := jwt_access.Subject
		ctx.Set("user_id", userID)
		return next(ctx)
	}
}
