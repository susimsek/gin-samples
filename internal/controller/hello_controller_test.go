package controller

import (
	"bytes"
	"gin-samples/internal/model"
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

func (m *MockHelloService) GetGreeting() model.Greeting {
	args := m.Called()
	return args.Get(0).(model.Greeting)
}

func (m *MockHelloService) CreateGreeting(input model.GreetingInput) (model.Greeting, error) {
	args := m.Called(input)
	return args.Get(0).(model.Greeting), args.Error(1)
}

func (m *MockHelloService) GetAllGreetings() []model.Greeting {
	args := m.Called()
	return args.Get(0).([]model.Greeting)
}

func TestHelloController_Hello(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockHelloService)
	mockService.On("GetGreeting").Return(model.Greeting{Message: "Mock Hello"})

	controller := NewHelloController(mockService, nil, nil)

	router := gin.Default()
	router.GET("/api/hello", controller.Hello)

	req, _ := http.NewRequest("GET", "/api/hello", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "Mock Hello"}`, w.Body.String())

	mockService.AssertExpectations(t)
}

func TestHelloController_CreateGreeting(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockHelloService)
	mockService.On("CreateGreeting", model.GreetingInput{Message: "Hello, Test!"}).
		Return(model.Greeting{Message: "Hello, Test!"}, nil)

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
	assert.JSONEq(t, `{"message": "Hello, Test!"}`, w.Body.String())

	mockService.AssertExpectations(t)
}

func TestHelloController_GetAllGreetings(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockHelloService)
	mockService.On("GetAllGreetings").Return([]model.Greeting{
		{Message: "Mock Hello"},
		{Message: "Mock Hi"},
	})

	controller := NewHelloController(mockService, nil, nil)

	router := gin.Default()
	router.GET("/api/hello/all", controller.GetAllGreetings)

	req, _ := http.NewRequest("GET", "/api/hello/all", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	expectedResponse := `[{"message": "Mock Hello"}, {"message": "Mock Hi"}]`
	assert.JSONEq(t, expectedResponse, w.Body.String())

	mockService.AssertExpectations(t)
}
