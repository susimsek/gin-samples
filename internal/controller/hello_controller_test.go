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
	"time"
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

func (m *MockHelloService) GetGreetingByID(id uint) (dto.GreetingResponse, error) {
	args := m.Called(id)
	return args.Get(0).(dto.GreetingResponse), args.Error(1)
}

func TestHelloController_Hello(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Mock Service
	mockService := new(MockHelloService)
	mockService.On("GetGreeting").Return(dto.GreetingResponse{
		ID:        1,
		Message:   "Mock Hello",
		CreatedAt: time.Date(2025, 1, 5, 10, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2025, 1, 5, 10, 0, 0, 0, time.UTC),
	})

	// Controller Setup
	controller := NewHelloController(mockService, nil, nil)
	router := gin.Default()
	router.GET("/api/hello", controller.Hello)

	// Mock Request
	req, _ := http.NewRequest("GET", "/api/hello", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	expectedResponse := `{
		"id": 1,
		"message": "Mock Hello",
		"createdAt": "2025-01-05T10:00:00Z",
		"updatedAt": "2025-01-05T10:00:00Z"
	}`
	assert.JSONEq(t, expectedResponse, w.Body.String())

	mockService.AssertExpectations(t)
}

func TestHelloController_CreateGreeting(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Mock Service
	mockService := new(MockHelloService)
	mockService.On("CreateGreeting", dto.GreetingInput{Message: "Hello, Test!"}).
		Return(dto.GreetingResponse{
			ID:        1,
			Message:   "Hello, Test!",
			CreatedAt: time.Date(2025, 1, 5, 10, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2025, 1, 5, 10, 0, 0, 0, time.UTC),
		}, nil)

	// Validator and Controller Setup
	validate := validator.New()
	controller := NewHelloController(mockService, validate, nil)
	router := gin.Default()
	router.POST("/api/hello", controller.CreateGreeting)

	// Mock Request
	body := []byte(`{"message": "Hello, Test!"}`)
	req, _ := http.NewRequest("POST", "/api/hello", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)

	expectedResponse := `{
		"id": 1,
		"message": "Hello, Test!",
		"createdAt": "2025-01-05T10:00:00Z",
		"updatedAt": "2025-01-05T10:00:00Z"
	}`
	assert.JSONEq(t, expectedResponse, w.Body.String())

	mockService.AssertExpectations(t)
}

func TestHelloController_GetAllGreetings(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Mock Service
	mockService := new(MockHelloService)
	mockService.On("GetAllGreetings").Return([]dto.GreetingResponse{
		{
			ID:        1,
			Message:   "Mock Hello",
			CreatedAt: time.Date(2025, 1, 5, 10, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2025, 1, 5, 10, 0, 0, 0, time.UTC),
		},
		{
			ID:        2,
			Message:   "Mock Hi",
			CreatedAt: time.Date(2025, 1, 5, 10, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2025, 1, 5, 10, 0, 0, 0, time.UTC),
		},
	}, nil)

	// Controller Setup
	controller := NewHelloController(mockService, nil, nil)
	router := gin.Default()
	router.GET("/api/hello/all", controller.GetAllGreetings)

	// Mock Request
	req, _ := http.NewRequest("GET", "/api/hello/all", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	expectedResponse := `[
		{
			"id": 1,
			"message": "Mock Hello",
			"createdAt": "2025-01-05T10:00:00Z",
			"updatedAt": "2025-01-05T10:00:00Z"
		},
		{
			"id": 2,
			"message": "Mock Hi",
			"createdAt": "2025-01-05T10:00:00Z",
			"updatedAt": "2025-01-05T10:00:00Z"
		}
	]`
	assert.JSONEq(t, expectedResponse, w.Body.String())

	mockService.AssertExpectations(t)
}

func TestHelloController_GetGreetingByID_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Mock Service
	mockService := new(MockHelloService)
	mockService.On("GetGreetingByID", uint(1)).Return(dto.GreetingResponse{
		ID:        1,
		Message:   "Mock Greeting",
		CreatedAt: time.Date(2025, 1, 5, 10, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2025, 1, 5, 10, 0, 0, 0, time.UTC),
	}, nil)

	// Controller Setup
	controller := NewHelloController(mockService, nil, nil)
	router := gin.Default()
	router.GET("/api/hello/:id", controller.GetGreetingByID)

	// Mock Request
	req, _ := http.NewRequest("GET", "/api/hello/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	expectedResponse := `{
		"id": 1,
		"message": "Mock Greeting",
		"createdAt": "2025-01-05T10:00:00Z",
		"updatedAt": "2025-01-05T10:00:00Z"
	}`
	assert.JSONEq(t, expectedResponse, w.Body.String())

	mockService.AssertExpectations(t)
}
