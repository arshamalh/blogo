package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/arshamalh/blogo/databases/mock"
	"github.com/arshamalh/blogo/models"
	"github.com/arshamalh/blogo/tools"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestCreatePost(t *testing.T) {
	mockDB := &mock.MockDatabase{}
	mockLogger, _ := zap.NewDevelopment()
	pc := NewPostController(mockDB, mockLogger)
	mockUserID := "mockUserID"
	mockUserContext := tools.CreateMockUserContext(mockUserID)
	postRequest := PostRequest{
		Title:      "Test Post",
		Content:    "This is a test post.",
		Categories: []string{"Test Category"},
	}

	jsonPostRequest, err := json.Marshal(postRequest)
	assert.NoError(t, err)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/posts", strings.NewReader(string(jsonPostRequest)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetRequest(req.WithContext(mockUserContext))

	err = pc.CreatePost(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Len(t, mockDB.GetPosts(), 1)
	var response map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Post created successfully", response["message"])
	assert.NotNil(t, response["post_id"])
}
