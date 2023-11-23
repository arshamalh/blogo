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
	"github.com/uptrace/bun"
	"go.uber.org/zap"
	"golang.org/x/net/context"
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
			Gl:     logger,
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
		FirstName: user.FirstName,
		LastName:  user.LastName,
		RoleID:    3,
	}
	if err := new_user.SetPassword(user.Password); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": "cannot register you"})
	}
	if _, err := uc.db.Insert(&new_user).Exec(context.Background()); err != nil {
		log.Gl.Error(err.Error())
		return ctx.JSON(http.StatusConflict, echo.Map{"message": "Failed to create user"})
	}
	log.Gl.Info("User created", zap.String("username", new_user.Username))
	return ctx.JSON(http.StatusCreated, echo.Map{"message": "user created", "uid": new_user.ID})
}

// CheckUsername checks the availability of a username.
// swagger:route GET /users/check-username CheckUsername
// Checks the availability of a username.
// responses:
//
//	200: map[string]interface{} "Username available"
//	400: map[string]interface{} "Invalid request"
//	409: map[string]interface{} "Username already taken"
func (uc *userController) UserLogin(ctx echo.Context) error {
	// Decode the body of request
	var user UserLoginRequest
	if ctx.Bind(&user) != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "invalid request"})
	}

	// Check if user exists
	var dbUser models.User
	if err := uc.db.NewSelect().Model(&dbUser).Where(bun.Where("username = ?", user.Username)).Scan(context.Background()); err != nil {
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
	return ctx.JSON(http.StatusOK, echo.Map{"user_id": ctx.Get("user_id")})
}
