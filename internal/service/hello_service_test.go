package service

import (
	customError "gin-samples/internal/error"
	"gin-samples/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockHelloRepository is a simplified mock implementation of HelloRepository
type MockHelloRepository struct{}

func (m *MockHelloRepository) ExistsByMessage(message string) bool {
	// Mock behavior for ExistsByMessage
	return message == "Duplicate Greeting"
}

func (m *MockHelloRepository) SaveGreeting(input model.GreetingInput) model.Greeting {
	// Mock behavior for SaveGreeting
	return model.Greeting(input)
}

func (m *MockHelloRepository) GetAllGreetings() []model.Greeting {
	// Mock behavior for GetAllGreetings
	return []model.Greeting{
		{Message: "Hello, World!"},
		{Message: "Hi there!"},
	}
}

func (m *MockHelloRepository) FindByMessage(message string) (*model.Greeting, bool) {
	// Mock behavior for FindByMessage
	if message == "Hello" {
		return &model.Greeting{Message: "Hello"}, true
	}
	return nil, false
}

func TestHelloService_GetGreeting(t *testing.T) {
	repo := &MockHelloRepository{}
	service := NewHelloService(repo)

	expected := model.Greeting{Message: "Hello, World!"}
	actual := service.GetGreeting()

	assert.Equal(t, expected, actual, "Greeting message should match the expected value")
}

func TestHelloService_CreateGreeting_Success(t *testing.T) {
	repo := &MockHelloRepository{}
	service := NewHelloService(repo)

	input := model.GreetingInput{Message: "Unique Greeting"}
	expected := model.Greeting{Message: "Unique Greeting"}

	actual, err := service.CreateGreeting(input)

	assert.NoError(t, err, "There should be no error")
	assert.Equal(t, expected, actual, "Created greeting should match the input message")
}

func TestHelloService_CreateGreeting_Conflict(t *testing.T) {
	repo := &MockHelloRepository{}
	service := NewHelloService(repo)

	input := model.GreetingInput{Message: "Duplicate Greeting"}

	_, err := service.CreateGreeting(input)

	assert.Error(t, err, "An error should be returned for duplicate message")
	assert.IsType(t, &customError.ResourceConflictError{}, err, "Error should be of type ResourceConflictError")

	conflictErr := err.(*customError.ResourceConflictError)
	assert.Equal(t, "Greeting", conflictErr.Resource, "Resource should be 'Greeting'")
	assert.Equal(t, "message", conflictErr.Criteria, "Criteria should be 'message'")
	assert.Equal(t, "Duplicate Greeting", conflictErr.Value, "Value should match the duplicate message")
}

func TestHelloService_GetAllGreetings(t *testing.T) {
	repo := &MockHelloRepository{}
	service := NewHelloService(repo)

	expected := []model.Greeting{
		{Message: "Hello, World!"},
		{Message: "Hi there!"},
	}

	actual := service.GetAllGreetings()

	assert.Equal(t, expected, actual, "All greetings should match the expected list")
}
