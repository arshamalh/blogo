// user.go
package controllers

import (
	"net/http"
	"time"

	"github.com/arshamalh/blogo/databases"
	"github.com/arshamalh/blogo/log"
	"github.com/arshamalh/blogo/models"
	"github.com/arshamalh/blogo/session"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// userController handles HTTP requests related to user management
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
	Username  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
	FirstName string `form:"first_name" json:"first_name"`
	LastName  string `form:"last_name" json:"last_name"`
}

// UserLoginRequest represents the request format for user login.
// swagger:parameters UserLogin
type UserLoginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// @Summary Register a user
// @Description Register a user
// @Accept json
// @Produce json
// @Param Username body string true "Username"
// @Param Password body string true "Password"
// @Param Email body string true "Email"
// @Param FirstName body string false "FirstName"
// @Param LastName body string false "LastName"
// @Success 201 {object} map[string]any "User created"
// @Failure 400 {object} map[string]any "Invalid request"
// @Failure 409 {object} map[string]any "Username already taken"
// @Router /users/register [post]
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
		FirstName: user.FirstName,
		LastName:  user.LastName,
		RoleID:    3,
	}
	if err := new_user.SetPassword(user.Password); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": "cannot register you"})
	}
	uid, err := uc.db.CreateUser(&new_user)
	if err != nil {
		log.Gl.Error(err.Error())
		return ctx.JSON(http.StatusConflict, echo.Map{"message": "Failed to create user"})
	}
	log.Gl.Info("User created", zap.String("username", new_user.Username))
	return ctx.JSON(http.StatusCreated, echo.Map{"message": "user created", "uid": uid})
}

// @Summary Login a user
// @Description Login a user
// @Accept json
// @Produce json
// @Param Username body string true "Username"
// @Param Password body string true "Password"
// @Success 200 {object} map[string]interface{} "Login success"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 401 {object} map[string]interface{} "Wrong password"
// @Failure 500 {object} map[string]interface{} "Error getting user"
// @Router /users/login [post]
func (uc *userController) UserLogin(ctx echo.Context) error {
	// Decode the body of request
	var user UserLoginRequest
	if ctx.Bind(&user) != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "invalid request"})
	}

	// Check if user exists
	dbUser, err := uc.db.GetUserByUsername(user.Username)
	if err != nil {
		log.Gl.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": "error getting user"})
	}

	// Check if password is correct
	if dbUser.ComparePasswords(user.Password) != nil {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"message": "wrong password"})
	}

	// Store username in the session
	sn := session.Create(dbUser.ID)

	// Generate access token and refresh token
	return ctx.JSON(http.StatusOK, echo.Map{"message": "login success", "session": sn})
}

// @Summary Check the availability of a username
// @Description Check the availability of a username
// @Accept json
// @Produce json
// @Param username body string true "Username to check"
// @Success 200 {object} map[string]interface{} "Username available"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 409 {object} map[string]interface{} "Username already taken"
// @Router /users/check-username [get]
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

// @Summary Logout a user
// @Description Logout a user
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Logout success"
// @Router /users/logout [post]
func (uc *userController) UserLogout(ctx echo.Context) error {
	ctx.SetCookie(&http.Cookie{
		Name:    "access_token",
		Path:    "/",
		Expires: time.Now(),
	})
	return ctx.JSON(http.StatusOK, echo.Map{"message": "logout success"})
}

// @Summary Retrieve the ID of the logged-in user
// @Description Retrieve the ID of the logged-in user
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "User ID retrieved"
// @Router /users/user-id [get]
func (uc *userController) UserID(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{"user_id": ctx.Get("user_id")})
}
