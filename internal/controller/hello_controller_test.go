package controller

import (
	"bytes"
	"gin-samples/internal/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockHelloService struct{}

func (m *MockHelloService) GetGreeting() model.Greeting {
	return model.Greeting{Message: "Mock Hello"}
}

func (m *MockHelloService) CreateGreeting(input model.GreetingInput) model.Greeting {
	return model.Greeting{Message: input.Message}
}

func TestHelloController(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockHelloService{}
	controller := NewHelloController(mockService)

	router := gin.Default()
	router.GET("/api/hello", controller.Hello)

	req, _ := http.NewRequest("GET", "/api/hello", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "Mock Hello"}`, w.Body.String())
}

func TestHelloController_CreateGreeting(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockHelloService{}
	controller := NewHelloController(mockService)

	router := gin.Default()
	router.POST("/api/hello", controller.CreateGreeting)

	// Test Data
	body := []byte(`{"message": "Hello, Test!"}`)
	req, _ := http.NewRequest("POST", "/api/hello", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, `{"message": "Hello, Test!"}`, w.Body.String())
}
