package controllers

import (
	"net/http"

	"github.com/arshamalh/blogo/database"
	"github.com/arshamalh/blogo/models"
	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
}

func UserRegister(c *gin.Context) {
	var user UserRequest
	if c.BindJSON(&user) == nil {
		if !database.CheckUserExists(user.Username) {
			new_user := models.User{
				Username: user.Username,
				Email:    user.Email,
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
	// Not implemented
}

func UserLogout(c *gin.Context) {
	// Not implemented
}
