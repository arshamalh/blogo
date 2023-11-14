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

var Gl *zap.Logger

func InitializeLogger() {
	cfg := zap.NewProductionConfig()
	logger, err := cfg.Build()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	defer logger.Sync()
	Gl = logger
}

type userController struct {
	basicAttributes
}

type UserRegisterRequest struct {
	Username  string `json:"username" form:"username" binding:"required"`
	Password  string `json:"password" form:"password" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

type UserLoginRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func NewUserController(db databases.Database, logger *zap.Logger) *userController {
	return &userController{
		basicAttributes: basicAttributes{
			db:     db,
			logger: logger,
		},
	}
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
	new_user.SetPassword(user.Password)
	uid, err := uc.db.CreateUser(&new_user)
	if err != nil {
		log.Gl.Error(err.Error())
		return ctx.JSON(http.StatusConflict, echo.Map{"message": "Failed to create user"})
	}
	log.Gl.Info("", zap.String("username", new_user.Username))
	return ctx.JSON(http.StatusCreated, echo.Map{"message": "user created", "uid": uid})
}

func (uc *userController) UserLogin(ctx echo.Context) error {
	// Decode the body of request
	var user UserLoginRequest
	if ctx.Bind(&user) != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "invalid request"})
	}

	// Check if user exists
	dbUser, err := uc.db.GetUserByUsername(ctx.FormValue("username"))
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
	log.Gl.Error(err.Error())
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

func (uc *userController) UserLogout(ctx echo.Context) error {
	ctx.SetCookie(&http.Cookie{
		Name:    "access_token",
		Path:    "/",
		Expires: time.Now(),
	})
	return ctx.JSON(http.StatusOK, echo.Map{"message": "logout success"})
}

func (uc *userController) UserID(ctx echo.Context) error {

	value := ctx.Get("user_id")
	return ctx.JSON(200, echo.Map{"user_id": value})
}
