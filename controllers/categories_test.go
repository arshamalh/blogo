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

func TestCreateCategory(t *testing.T) {
	mockDB := &mock.MockDatabase{}
	mockLogger, _ := zap.NewDevelopment()
	cc := NewCategoryController(mockDB, mockLogger)
	category := models.Category{
		Name: "TestCategory",
	}

	jsonCategory, err := json.Marshal(category)
	assert.NoError(t, err)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/categories", strings.NewReader(string(jsonCategory)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err = cc.CreateCategory(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.True(t, mockDB.CheckCategoryExists("TestCategory"))
}
