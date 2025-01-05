package service

import (
	"errors"
	"gin-samples/internal/dto"
	"gin-samples/internal/entity"
	customError "gin-samples/internal/error"
	customMock "gin-samples/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

// MockHelloRepository is a mock implementation of HelloRepository
type MockHelloRepository struct {
	mock.Mock
}

func (m *MockHelloRepository) ExistsByMessage(message string) (bool, error) {
	args := m.Called(message)
	return args.Bool(0), args.Error(1)
}

func (m *MockHelloRepository) SaveGreeting(greeting *entity.Greeting) (*entity.Greeting, error) {
	args := m.Called(greeting)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Greeting), args.Error(1)
}

func (m *MockHelloRepository) GetAllGreetings() ([]entity.Greeting, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Greeting), args.Error(1)
}

func (m *MockHelloRepository) FindByMessage(message string) (*entity.Greeting, error) {
	args := m.Called(message)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Greeting), args.Error(1)
}

// Test cases for HelloService

func TestHelloService_GetGreeting(t *testing.T) {
	mockClock := new(customMock.MockClock)
	fixedTime := time.Date(2025, 1, 5, 10, 0, 0, 0, time.UTC)
	mockClock.On("Now").Return(fixedTime)
	service := NewHelloService(nil, mockClock) // No repo needed for this method

	expected := dto.GreetingResponse{ID: 0,
		Message:   "Hello, World!",
		CreatedAt: fixedTime,
		UpdatedAt: fixedTime}
	actual := service.GetGreeting()

	assert.Equal(t, expected, actual, "Greeting message should match the expected value")
}

func TestHelloService_CreateGreeting_Success(t *testing.T) {
	mockRepo := new(MockHelloRepository)
	service := NewHelloService(mockRepo, nil)

	input := dto.GreetingInput{Message: "Unique Greeting"}
	expectedEntity := &entity.Greeting{ID: 1, Message: "Unique Greeting"}
	expectedResponse := dto.GreetingResponse{ID: 1, Message: "Unique Greeting"}

	mockRepo.On("ExistsByMessage", input.Message).Return(false, nil)
	mockRepo.On("SaveGreeting", mock.AnythingOfType("*entity.Greeting")).Return(expectedEntity, nil)

	actual, err := service.CreateGreeting(input)

	assert.NoError(t, err, "There should be no error")
	assert.Equal(t, expectedResponse, actual, "Created greeting should match the expected response")

	mockRepo.AssertExpectations(t)
}

func TestHelloService_CreateGreeting_Conflict(t *testing.T) {
	mockRepo := new(MockHelloRepository)
	service := NewHelloService(mockRepo, nil)

	input := dto.GreetingInput{Message: "Duplicate Greeting"}

	mockRepo.On("ExistsByMessage", input.Message).Return(true, nil)

	_, err := service.CreateGreeting(input)

	assert.Error(t, err, "An error should be returned for duplicate message")
	assert.IsType(t, &customError.ResourceConflictError{}, err, "Error should be of type ResourceConflictError")

	var conflictErr *customError.ResourceConflictError
	errors.As(err, &conflictErr)
	assert.Equal(t, "Greeting", conflictErr.Resource, "Resource should be 'Greeting'")
	assert.Equal(t, "message", conflictErr.Criteria, "Criteria should be 'message'")
	assert.Equal(t, "Duplicate Greeting", conflictErr.Value, "Value should match the duplicate message")

	mockRepo.AssertExpectations(t)
}

func TestHelloService_GetAllGreetings(t *testing.T) {
	mockRepo := new(MockHelloRepository)
	service := NewHelloService(mockRepo, nil)

	expectedEntities := []entity.Greeting{
		{ID: 1, Message: "Hello, World!"},
		{ID: 2, Message: "Hi there!"},
	}
	expectedResponses := []dto.GreetingResponse{
		{ID: 1, Message: "Hello, World!"},
		{ID: 2, Message: "Hi there!"},
	}

	mockRepo.On("GetAllGreetings").Return(expectedEntities, nil)

	actual, err := service.GetAllGreetings()

	assert.NoError(t, err, "There should be no error")
	assert.Equal(t, expectedResponses, actual, "All greetings should match the expected DTO list")

	mockRepo.AssertExpectations(t)
}
