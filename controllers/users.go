package controllers

import (
	"net/http"
	"time"

	"github.com/arshamalh/blogo/databases"
	"github.com/arshamalh/blogo/models"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

var Gl *zap.Logger

func init() {
	InitializeLogger()
}

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

func NewUserController(db databases.Database, logger *zap.Logger) *userController {
	return &userController{
		basicAttributes: basicAttributes{
			db:     db,
			logger: logger,
		},
	}
}

type UserRegisterRequest struct {
	Username  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
	FirstName string `form:"first_name" json:"first_name"`
	LastName  string `form:"last_name" json:"last_name"`
}

type UserLoginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

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
		Gl.Error("Failed to create user", zap.Error(err), zap.String("username", user.Username))
		return ctx.JSON(http.StatusConflict, echo.Map{"message": "Failed to create user"})
	}
	Gl.Info("User created", zap.String("username", new_user.Username))
	return ctx.JSON(http.StatusCreated, echo.Map{"message": "user created", "uid": uid})
}

func (uc *userController) UserLogin(ctx echo.Context) error {
	var user UserLoginRequest
	if ctx.Bind(&user) != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "invalid request"})
	}

	dbUser, err := uc.db.GetUserByUsername(ctx.FormValue("username"))
	if err != nil {
		Gl.Error("Error getting user", zap.Error(err))
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": "error getting user"})
	}

	if dbUser.ComparePasswords(user.Password) != nil {
		Gl.Error("Wrong password", zap.String("username", dbUser.Username))
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"message": "wrong password"})
	}
	sn := "session"
	Gl.Info("User logged in", zap.String("username", dbUser.Username))
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

	return ctx.JSON(http.StatusOK, echo.Map{"message": "logout success"})
}
