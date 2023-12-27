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

func TestCreateComment(t *testing.T) {
	mockDB := &mock.MockDatabase{}
	mockLogger, _ := zap.NewDevelopment()
	cc := NewCommentController(mockDB)
	mockUserID := "mockUserID"
	mockUserContext := tools.CreateMockUserContext(mockUserID)
	comment := models.Comment{}

	jsonComment, err := json.Marshal(comment)
	assert.NoError(t, err)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/comments", strings.NewReader(string(jsonComment)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetRequest(req.WithContext(mockUserContext))
	err = cc.CreateComment(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Len(t, mockDB.GetComments(), 1)
	var response map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, mockUserID, response["author_id"])
	assert.NotNil(t, response["comment"])
}
