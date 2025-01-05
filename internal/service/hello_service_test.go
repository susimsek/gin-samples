package service

import (
	"errors"
	"gin-samples/internal/domain"
	"gin-samples/internal/dto"
	customError "gin-samples/internal/error"
	customMock "gin-samples/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

// Test cases for HelloService

func TestHelloService_GetGreeting(t *testing.T) {
	mockClock := new(customMock.MockClock)
	fixedTime := time.Date(2025, 1, 5, 10, 0, 0, 0, time.UTC)
	mockClock.On("Now").Return(fixedTime)
	mockMapper := new(customMock.MockHelloMapper)

	service := NewHelloService(nil, mockMapper, mockClock) // No repo needed for this method

	expected := dto.GreetingResponse{
		ID:        0,
		Message:   "Hello, World!",
		CreatedAt: fixedTime,
		UpdatedAt: fixedTime,
	}
	actual := service.GetGreeting()

	assert.Equal(t, expected, actual, "Greeting message should match the expected value")
}

func TestHelloService_CreateGreeting_Success(t *testing.T) {
	mockRepo := new(customMock.MockHelloRepository)
	mockMapper := new(customMock.MockHelloMapper)
	mockClock := new(customMock.MockClock)

	input := dto.GreetingInput{Message: "Unique Greeting"}
	expectedEntity := domain.Greeting{ID: 1, Message: "Unique Greeting"}
	expectedResponse := dto.GreetingResponse{ID: 1, Message: "Unique Greeting"}

	mockRepo.On("ExistsByMessage", input.Message).Return(false, nil)
	mockRepo.On("SaveGreeting", mock.AnythingOfType("*domain.Greeting")).Return(&expectedEntity, nil)
	mockMapper.On("ToGreetingEntity", input).Return(expectedEntity, nil)
	mockMapper.On("ToGreetingResponse", expectedEntity).Return(expectedResponse, nil)

	service := NewHelloService(mockRepo, mockMapper, mockClock)

	actual, err := service.CreateGreeting(input)

	assert.NoError(t, err, "There should be no error")
	assert.Equal(t, expectedResponse, actual, "Created greeting should match the expected response")

	mockRepo.AssertExpectations(t)
	mockMapper.AssertExpectations(t)
}

func TestHelloService_CreateGreeting_Conflict(t *testing.T) {
	mockRepo := new(customMock.MockHelloRepository)
	mockMapper := new(customMock.MockHelloMapper)
	mockClock := new(customMock.MockClock)

	input := dto.GreetingInput{Message: "Duplicate Greeting"}

	mockRepo.On("ExistsByMessage", input.Message).Return(true, nil)

	service := NewHelloService(mockRepo, mockMapper, mockClock)

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
	mockRepo := new(customMock.MockHelloRepository)
	mockMapper := new(customMock.MockHelloMapper)
	mockClock := new(customMock.MockClock)

	expectedEntities := []domain.Greeting{
		{ID: 1, Message: "Hello, World!"},
		{ID: 2, Message: "Hi there!"},
	}
	expectedResponses := []dto.GreetingResponse{
		{ID: 1, Message: "Hello, World!"},
		{ID: 2, Message: "Hi there!"},
	}

	mockRepo.On("GetAllGreetings").Return(expectedEntities, nil)
	mockMapper.On("ToGreetingResponses", expectedEntities).Return(expectedResponses, nil)

	service := NewHelloService(mockRepo, mockMapper, mockClock)

	actual, err := service.GetAllGreetings()

	assert.NoError(t, err, "There should be no error")
	assert.Equal(t, expectedResponses, actual, "All greetings should match the expected DTO list")

	mockRepo.AssertExpectations(t)
	mockMapper.AssertExpectations(t)
}
