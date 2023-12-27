package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/arshamalh/blogo/databases/mock"
	"github.com/arshamalh/blogo/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestUserRegister(t *testing.T) {
	mockDB := &mock.MockDatabase{}
	mockLogger, _ := zap.NewDevelopment()
	uc := NewUserController(mockDB, mockLogger)
	userRequest := UserRegisterRequest{
		Username:  "testuser",
		Password:  "testpassword",
		Email:     "",
		FirstName: "John",
		LastName:  "Doe",
	}
	jsonUserRequest, err := json.Marshal(userRequest)
	assert.NoError(t, err)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users/register", strings.NewReader(string(jsonUserRequest)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err = uc.UserRegister(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	assert.Len(t, mockDB.GetUsers(), 1)

	var response map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "user created", response["message"])
	assert.NotNil(t, response["uid"])
}

func TestUserLogin(t *testing.T) {

	mockDB := &mock.MockDatabase{}
	mockLogger, _ := zap.NewDevelopment()
	uc := NewUserController(mockDB, mockLogger)

	userLoginRequest := UserLoginRequest{
		Username: "testuser",
		Password: "testpassword",
	}

	jsonUserLoginRequest, err := json.Marshal(userLoginRequest)
	assert.NoError(t, err)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(string(jsonUserLoginRequest)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err = uc.UserLogin(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	var response map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "login success", response["message"])
	assert.NotNil(t, response["session"])
}
