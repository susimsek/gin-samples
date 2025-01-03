package controller

import (
	"gin-samples/internal/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockHelloService struct{}

func (m *mockHelloService) GetGreeting() model.Greeting {
	return model.Greeting{Message: "Mock Hello"}
}

func TestHelloController(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &mockHelloService{}
	controller := NewHelloController(mockService)

	router := gin.Default()
	router.GET("/api/hello", controller.Hello)

	req, _ := http.NewRequest("GET", "/api/hello", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "Mock Hello"}`, w.Body.String())
}
