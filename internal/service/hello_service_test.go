package service

import (
	"errors"
	customError "gin-samples/internal/error"
	"gin-samples/internal/model"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockHelloRepository is a simplified mock implementation of HelloRepository
type MockHelloRepository struct {
	mock.Mock
}

func (m *MockHelloRepository) ExistsByMessage(message string) bool {
	args := m.Called(message)
	return args.Bool(0)
}

func (m *MockHelloRepository) SaveGreeting(input model.GreetingInput) model.Greeting {
	args := m.Called(input)
	return args.Get(0).(model.Greeting)
}

func (m *MockHelloRepository) GetAllGreetings() []model.Greeting {
	args := m.Called()
	return args.Get(0).([]model.Greeting)
}

func (m *MockHelloRepository) FindByMessage(message string) (*model.Greeting, bool) {
	args := m.Called(message)
	if args.Get(0) == nil {
		return nil, args.Bool(1)
	}
	return args.Get(0).(*model.Greeting), args.Bool(1)
}

func TestHelloService_GetGreeting(t *testing.T) {
	mockRepo := new(MockHelloRepository)
	service := NewHelloService(mockRepo)

	expected := model.Greeting{Message: "Hello, World!"}
	actual := service.GetGreeting()

	assert.Equal(t, expected, actual, "Greeting message should match the expected value")
}

func TestHelloService_CreateGreeting_Success(t *testing.T) {
	mockRepo := new(MockHelloRepository)
	service := NewHelloService(mockRepo)

	input := model.GreetingInput{Message: "Unique Greeting"}
	expected := model.Greeting{Message: "Unique Greeting"}

	mockRepo.On("ExistsByMessage", input.Message).Return(false)
	mockRepo.On("SaveGreeting", input).Return(expected)

	actual, err := service.CreateGreeting(input)

	assert.NoError(t, err, "There should be no error")
	assert.Equal(t, expected, actual, "Created greeting should match the input message")

	mockRepo.AssertExpectations(t)
}

func TestHelloService_CreateGreeting_Conflict(t *testing.T) {
	mockRepo := new(MockHelloRepository)
	service := NewHelloService(mockRepo)

	input := model.GreetingInput{Message: "Duplicate Greeting"}

	mockRepo.On("ExistsByMessage", input.Message).Return(true)

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
	service := NewHelloService(mockRepo)

	expected := []model.Greeting{
		{Message: "Hello, World!"},
		{Message: "Hi there!"},
	}

	mockRepo.On("GetAllGreetings").Return(expected)

	actual := service.GetAllGreetings()

	assert.Equal(t, expected, actual, "All greetings should match the expected list")

	mockRepo.AssertExpectations(t)
}
