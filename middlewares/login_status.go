package middlewares

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func IsLoggedIn(c *gin.Context) {
	jwt_token, _ := c.Cookie("jwt")
	token, err := jwt.ParseWithClaims(
		jwt_token, &jwt.StandardClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		},
	)
	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "you are not allowed to make new posts",
		})
	} else {
		payload := token.Claims.(*jwt.StandardClaims)
		id, _ := strconv.Atoi(payload.Subject)
		c.Set("user_id", id)
		c.Next()
	}
}
