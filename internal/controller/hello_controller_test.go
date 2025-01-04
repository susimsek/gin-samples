package controller

import (
	"bytes"
	"gin-samples/internal/model"
	"github.com/gin-gonic/gin"
	en "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockHelloService simulates the HelloService
type MockHelloService struct{}

func (m *MockHelloService) GetGreeting() model.Greeting {
	return model.Greeting{Message: "Mock Hello"}
}

func (m *MockHelloService) CreateGreeting(input model.GreetingInput) (model.Greeting, error) {
	return model.Greeting{Message: input.Message}, nil
}

func (m *MockHelloService) GetAllGreetings() []model.Greeting {
	return []model.Greeting{
		{Message: "Mock Hello"},
		{Message: "Mock Hi"},
	}
}

func TestHelloController_Hello(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockHelloService{}
	controller := NewHelloController(mockService, nil, nil)

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

	// Mock Validator ve Translator
	validate := validator.New()
	translator := en.New()
	uni := ut.New(translator, translator)
	trans, _ := uni.GetTranslator("en")

	mockService := &MockHelloService{}
	controller := NewHelloController(mockService, validate, trans)

	router := gin.Default()
	router.POST("/api/hello", controller.CreateGreeting)

	body := []byte(`{"message": "Hello, Test!"}`)
	req, _ := http.NewRequest("POST", "/api/hello", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, `{"message": "Hello, Test!"}`, w.Body.String())
}

func TestHelloController_GetAllGreetings(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockHelloService{}
	controller := NewHelloController(mockService, nil, nil)

	router := gin.Default()
	router.GET("/api/hello/all", controller.GetAllGreetings)

	req, _ := http.NewRequest("GET", "/api/hello/all", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	expectedResponse := `[{"message": "Mock Hello"}, {"message": "Mock Hi"}]`
	assert.JSONEq(t, expectedResponse, w.Body.String())
}
