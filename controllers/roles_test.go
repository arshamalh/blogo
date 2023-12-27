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

func TestCreateRole(t *testing.T) {
	mockDB := &mock.MockDatabase{}
	mockLogger, _ := zap.NewDevelopment()
	rc := NewRoleController(mockDB, mockLogger)
	roleRequest := models.Role{
		Name:        "TestRole",
		Description: "This is a test role.",
	}
	jsonRoleRequest, err := json.Marshal(roleRequest)
	assert.NoError(t, err)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/roles", strings.NewReader(string(jsonRoleRequest)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err = rc.CreateRole(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Len(t, mockDB.GetRoles(), 1)
	var response map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotNil(t, response["role"])
}
