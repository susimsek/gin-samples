package router_test

import (
	"gin-samples/internal/mock"
	"gin-samples/internal/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	en "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/stretchr/testify/assert"
)

func TestSetupRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockController := &mock.MockHelloController{}
	mockHealthController := &mock.MockHealthController{}
	translator := en.New()
	uni := ut.New(translator, translator)
	trans, _ := uni.GetTranslator("en")

	r := router.SetupRouter(mockController, mockHealthController, trans)

	req, _ := http.NewRequest(http.MethodGet, "/api/hello", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Adjust expected response to match actual response
	expectedResponse := `{
		"id": 1,
		"message": "Mocked Hello, World!",
		"createdAt": "2025-01-05T10:00:00Z",
		"updatedAt": "2025-01-05T10:00:00Z"
	}`
	assert.JSONEq(t, expectedResponse, w.Body.String())
}
