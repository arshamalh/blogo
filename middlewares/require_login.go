package middlewares

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/arshamalh/blogo/session"
	"github.com/arshamalh/blogo/tools"
	"github.com/gin-gonic/gin"
)

func RequireLogin(ctx *gin.Context) {
	access_token, _ := ctx.Cookie("access_token")
	refresh_token, _ := ctx.Cookie("refresh_token")

	if access_token != "" {
		jwt_access, err := tools.ExtractTokenData(access_token, os.Getenv("JWT_SECRET"))

		// If access token is valid and not expired, extract data from it
		if err == nil {
			ctx.Set("user_id", jwt_access.Subject)
			ctx.Next()
			return
		}
	}

	if refresh_token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "you are not allowed to access this route",
		})
		return
	}

	// If access token is invalid, check if refresh token is valid and make new access token
	jwt_refresh, err := tools.ExtractTokenData(refresh_token, os.Getenv("REFRESH_TOKEN_SECRET"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "you are not allowed to access this route",
		})
		return
	}

	// Get id from session, and create new access token	and refresh token
	// Browser automatically deletes old access token, so we should get user id from session
	sn := session.Get(jwt_refresh.Subject)
	user_id := strconv.Itoa(int(sn.UserID))
	access_token, _ = tools.GenerateToken(user_id, time.Hour*1, os.Getenv("JWT_SECRET"))
	refresh_token, _ = tools.GenerateToken(jwt_refresh.Subject, time.Hour*24*7, os.Getenv("REFRESH_TOKEN_SECRET"))
	ctx.SetCookie("access_token", access_token, 10, "/", "", false, true)
	ctx.SetCookie("refresh_token", refresh_token, 50, "/", "", false, true)
	ctx.Set("user_id", user_id)
	ctx.Next()
}
