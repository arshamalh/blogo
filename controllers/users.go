package controllers

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/arshamalh/blogo/database"
	"github.com/arshamalh/blogo/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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
				payload := jwt.StandardClaims{
					Subject:   strconv.Itoa(int(db_user.ID)),
					ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
				}
				token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(os.Getenv("JWT_SECRET")))
				c.SetCookie("jwt", token, 86400, "/", "", false, true)
				c.JSON(http.StatusOK, gin.H{"status": "login success", "token": token})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "wrong password"})
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "user not found"})
		}
	}
}

func UserLogout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "logout success"})
}

func UserID(c *gin.Context) {
	value, _ := c.Get("user_id")
	c.JSON(200, gin.H{"user_id": value})
}
