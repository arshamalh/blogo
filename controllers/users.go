package controllers

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/arshamalh/blogo/database"
	"github.com/arshamalh/blogo/models"
	"github.com/arshamalh/blogo/tools"
	"github.com/gin-gonic/gin"
)

type UserRegisterRequest struct {
	Username  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
	FisrtName string `form:"first_name" json:"first_name"`
	LastName  string `form:"last_name" json:"last_name"`
}

type UserLoginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func UserRegister(c *gin.Context) {
	var user UserRegisterRequest
	if c.BindJSON(&user) == nil {
		if !database.CheckUserExists(user.Username) {
			new_user := models.User{
				Username:  user.Username,
				Email:     user.Email,
				FisrtName: user.FisrtName,
				LastName:  user.LastName,
			}
			new_user.SetPassword(user.Password)
			database.CreateUser(&new_user)
			c.JSON(http.StatusOK, gin.H{"status": "user created"})
		} else {
			c.JSON(http.StatusConflict, gin.H{"status": "user already exists"})
		}
	}
}

func CheckUsername(c *gin.Context) {
	var username string
	if c.BindJSON(&username) == nil {
		if !database.CheckUserExists(username) {
			c.JSON(http.StatusOK, gin.H{"status": "username available"})
		} else {
			c.JSON(http.StatusConflict, gin.H{"status": "username has already taken"})
		}
	}
}

func UserLogin(c *gin.Context) {
	var user UserLoginRequest
	if c.BindJSON(&user) == nil {
		if db_user, _ := database.GetUserByUsername(user.Username); db_user.ID != 0 {
			if db_user.ComparePasswords(user.Password) == nil {
				db_user_id := strconv.Itoa(int(db_user.ID))
				access_token, _ := tools.GenerateToken(db_user_id, time.Hour*1, os.Getenv("JWT_SECRET"))
				refresh_token, _ := tools.GenerateToken(db_user_id, time.Hour*24*7, os.Getenv("REFRESH_TOKEN_SECRET"))
				c.SetCookie("access_token", access_token, 3600, "/", "", false, true)
				c.SetCookie("refresh_token", refresh_token, 3600*24*7, "/refresh_token", "", false, true)
				c.JSON(http.StatusOK, gin.H{"status": "login success"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "wrong password"})
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "user not found"})
		}
	}
}

// This route should include refresh token and will return new access token
func RefreshToken(c *gin.Context) {
	refresh_token, err := c.Cookie("refresh_token")
	if err != nil || refresh_token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "you are logged out"})
		return
	}
	jwt, err := tools.ExtractTokenData(refresh_token, os.Getenv("REFRESH_TOKEN_SECRET"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "you are logged out"})
		return
	}
	access_token, _ := tools.GenerateToken(jwt.Subject, time.Hour*1, os.Getenv("JWT_SECRET"))
	refresh_token, _ = tools.GenerateToken(jwt.Subject, time.Hour*24*7, os.Getenv("REFRESH_TOKEN_SECRET"))
	c.SetCookie("access_token", access_token, 3600, "/", "", false, true)
	c.SetCookie("refresh_token", refresh_token, 3600*24*7, "/refresh_token", "", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "token refreshed"})
}

func UserLogout(c *gin.Context) {
	c.SetCookie("refresh_token", "", -1, "/refresh_token", "", false, true)
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "logout success"})
}

func UserID(c *gin.Context) {
	value, _ := c.Get("user_id")
	c.JSON(200, gin.H{"user_id": value})
}
