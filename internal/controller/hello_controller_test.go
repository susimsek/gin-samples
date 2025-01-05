package controller

import (
	"bytes"
	"gin-samples/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockHelloService simulates the HelloService
type MockHelloService struct {
	mock.Mock
}

func (m *MockHelloService) GetGreeting() dto.GreetingResponse {
	args := m.Called()
	return args.Get(0).(dto.GreetingResponse)
}

func (m *MockHelloService) CreateGreeting(input dto.GreetingInput) (dto.GreetingResponse, error) {
	args := m.Called(input)
	return args.Get(0).(dto.GreetingResponse), args.Error(1)
}

func (m *MockHelloService) GetAllGreetings() ([]dto.GreetingResponse, error) {
	args := m.Called()
	return args.Get(0).([]dto.GreetingResponse), args.Error(1)
}

func TestHelloController_Hello(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockHelloService)
	mockService.On("GetGreeting").Return(dto.GreetingResponse{ID: 1, Message: "Mock Hello"})

	controller := NewHelloController(mockService, nil, nil)

	router := gin.Default()
	router.GET("/api/hello", controller.Hello)

	req, _ := http.NewRequest("GET", "/api/hello", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"id": 1, "message": "Mock Hello"}`, w.Body.String())

	mockService.AssertExpectations(t)
}

func TestHelloController_CreateGreeting(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockHelloService)
	mockService.On("CreateGreeting", dto.GreetingInput{Message: "Hello, Test!"}).
		Return(dto.GreetingResponse{ID: 1, Message: "Hello, Test!"}, nil)

	validate := validator.New()
	controller := NewHelloController(mockService, validate, nil)

	router := gin.Default()
	router.POST("/api/hello", controller.CreateGreeting)

	body := []byte(`{"message": "Hello, Test!"}`)
	req, _ := http.NewRequest("POST", "/api/hello", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, `{"id": 1, "message": "Hello, Test!"}`, w.Body.String())

	mockService.AssertExpectations(t)
}

func TestHelloController_GetAllGreetings(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockHelloService)
	mockService.On("GetAllGreetings").Return([]dto.GreetingResponse{
		{ID: 1, Message: "Mock Hello"},
		{ID: 2, Message: "Mock Hi"},
	}, nil)

	controller := NewHelloController(mockService, nil, nil)

	router := gin.Default()
	router.GET("/api/hello/all", controller.GetAllGreetings)

	req, _ := http.NewRequest("GET", "/api/hello/all", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	expectedResponse := `[{"id": 1, "message": "Mock Hello"}, {"id": 2, "message": "Mock Hi"}]`
	assert.JSONEq(t, expectedResponse, w.Body.String())

	mockService.AssertExpectations(t)
}
