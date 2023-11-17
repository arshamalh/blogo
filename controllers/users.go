// user.go
package controllers

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/arshamalh/blogo/databases"
	"github.com/arshamalh/blogo/models"
	"github.com/arshamalh/blogo/session"
	"github.com/arshamalh/blogo/tools"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// userController handles HTTP requests related to user management.
type userController struct {
	basicAttributes
}

// NewUserController creates a new instance of userController.
func NewUserController(db databases.Database, logger *zap.Logger) *userController {
	return &userController{
		basicAttributes: basicAttributes{
			db:     db,
			logger: logger,
		},
	}
}

// UserRegisterRequest represents the request format for user registration.
// swagger:parameters UserRegister
type UserRegisterRequest struct {
	// in: body
	Username  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
	FisrtName string `form:"first_name" json:"first_name"`
	LastName  string `form:"last_name" json:"last_name"`
}

// UserLoginRequest represents the request format for user login.
// swagger:parameters UserLogin
type UserLoginRequest struct {
	// in: body
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// ShowAccount   godoc
// @Summary      Register a user
// @Description  Register a user
// @Accept       json
// @Produce      json
// @Param        Username body string true "Username"
// @Param        Password body string true "Password"
// @Param        Email body string true "Email"
// @Param        FirstName body string false "FirstName"
// @Param        LastName body string false "LastName"
// @Success      201 {object} map[string]any
// @Failure      400 {object} map[string]any
// @Failure      409 {object} map[string]any
// @Router       /users/register [post]
func (uc *userController) UserRegister(ctx echo.Context) error {
	var user UserRegisterRequest
	if ctx.Bind(&user) != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "invalid request"})
	}
	if uc.db.CheckUserExists(user.Username) {
		return ctx.JSON(http.StatusConflict, echo.Map{"message": "user already exists"})
	}
	new_user := models.User{
		Username:  user.Username,
		Email:     user.Email,
		FisrtName: user.FisrtName,
		LastName:  user.LastName,
		RoleID:    3,
	}
	new_user.SetPassword(user.Password)
	uid, err := uc.db.CreateUser(&new_user)
	if err != nil {
		return ctx.JSON(http.StatusConflict, err)
	}
	return ctx.JSON(http.StatusCreated, echo.Map{"message": "user created", "uid": uid})
}

// CheckUsername checks the availability of a username.
// swagger:route GET /users/check-username CheckUsername
// Checks the availability of a username.
// responses:
//
//	200: map[string]interface{} "Username available"
//	400: map[string]interface{} "Invalid request"
//	409: map[string]interface{} "Username already taken"
func (uc *userController) CheckUsername(ctx echo.Context) error {
	var username string
	if ctx.Bind(&username) != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "invalid request"})
	} else if uc.db.CheckUserExists(username) {
		return ctx.JSON(http.StatusConflict, echo.Map{"message": "username is already taken"})
	} else {
		return ctx.JSON(http.StatusOK, echo.Map{"message": "username available"})
	}
}

// UserLogin logs in a user.
// swagger:route POST /users/login UserLogin
// Logs in a user.
// responses:
//
//	200: map[string]interface{} "Login success"
//	400: map[string]interface{} "Invalid request"
//	401: map[string]interface{} "User not found"
//	500: map[string]interface{} "Internal server error"
func (uc *userController) UserLogin(ctx echo.Context) error {
	// Decode the body of request
	var user UserLoginRequest
	if ctx.Bind(&user) != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "invalid request"})

	}

	// Check if user exists
	db_user, err := uc.db.GetUserByUsername(user.Username)
	if err != nil {
		if db_user.ID == 0 {
			return ctx.JSON(http.StatusUnauthorized, echo.Map{"message": "user not found"})

		}
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": "error getting user"})

	}

	// Check if password is correct
	if db_user.ComparePasswords(user.Password) != nil {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"message": "wrong password"})
	}

	// Store username in the session
	sn := session.Create(db_user.ID)

	// Generate access token and refresh token
	access_token, _ := tools.GenerateToken(strconv.Itoa(int(db_user.ID)), time.Hour*1, os.Getenv("JWT_SECRET"))
	ctx.SetCookie(&http.Cookie{
		Name:     "access_token",
		Path:     "/",
		Value:    access_token,
		Expires:  time.Now().Add(time.Hour * 24 * 7),
		HttpOnly: true,
	})
	return ctx.JSON(http.StatusOK, echo.Map{"message": "login success", "session": sn})
}

// UserLogout logs out a user.
// swagger:route POST /users/logout UserLogout
// Logs out a user.
// responses:
//
//	200: map[string]interface{} "Logout success"
func (uc *userController) UserLogout(ctx echo.Context) error {
	ctx.SetCookie(&http.Cookie{
		Name:    "access_token",
		Path:    "/",
		Expires: time.Now(),
	})
	return ctx.JSON(http.StatusOK, echo.Map{"message": "logout success"})
}

// UserID retrieves the ID of the logged-in user.
// swagger:route GET /users/user-id UserID
// Retrieves the ID of the logged-in user.
// responses:
//
//	200: map[string]interface{} "User ID retrieved"
func (uc *userController) UserID(ctx echo.Context) error {
	value := ctx.Get("user_id")
	return ctx.JSON(200, echo.Map{"user_id": value})
}
