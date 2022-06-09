package middlewares

import (
	"net/http"
	"os"

	"github.com/arshamalh/blogo/tools"
	"github.com/gin-gonic/gin"
)

func IsLoggedIn(c *gin.Context) {
	access_token, _ := c.Cookie("access_token")
	jwt, err := tools.ExtractTokenData(access_token, os.Getenv("JWT_SECRET"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "you are not allowed to access this route",
		})
		return
	}
	c.Set("user_id", jwt.Subject)
	c.Next()
}
