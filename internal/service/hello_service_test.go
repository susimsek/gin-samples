package service

import (
	"errors"
	"gin-samples/internal/domain"
	"gin-samples/internal/dto"
	customError "gin-samples/internal/error"
	customMock "gin-samples/internal/mock"
	"gin-samples/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

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
	mockRepo.On("Save", mock.AnythingOfType("domain.Greeting")).Return(expectedEntity, nil)
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

	mockRepo.On("FindAll").Return(expectedEntities, nil)
	mockMapper.On("ToGreetingResponses", expectedEntities).Return(expectedResponses, nil)

	service := NewHelloService(mockRepo, mockMapper, mockClock)

	actual, err := service.GetAllGreetings()

	assert.NoError(t, err, "There should be no error")
	assert.Equal(t, expectedResponses, actual, "All greetings should match the expected DTO list")

	mockRepo.AssertExpectations(t)
	mockMapper.AssertExpectations(t)
}

func TestHelloService_GetGreetingByID_Success(t *testing.T) {
	mockRepo := new(customMock.MockHelloRepository)
	mockMapper := new(customMock.MockHelloMapper)
	mockClock := new(customMock.MockClock)

	expectedEntity := domain.Greeting{ID: 1, Message: "Hello, Mock!"}
	expectedResponse := dto.GreetingResponse{ID: 1, Message: "Hello, Mock!"}

	mockRepo.On("FindByID", uint(1)).Return(util.Optional[domain.Greeting]{Value: &expectedEntity}, nil)
	mockMapper.On("ToGreetingResponse", expectedEntity).Return(expectedResponse, nil)

	service := NewHelloService(mockRepo, mockMapper, mockClock)

	actual, err := service.GetGreetingByID(1)

	assert.NoError(t, err, "There should be no error")
	assert.Equal(t, expectedResponse, actual, "Greeting should match the expected response")

	mockRepo.AssertExpectations(t)
	mockMapper.AssertExpectations(t)
}

func TestHelloService_GetGreetingByID_NotFound(t *testing.T) {
	mockRepo := new(customMock.MockHelloRepository)
	mockMapper := new(customMock.MockHelloMapper)
	mockClock := new(customMock.MockClock)

	mockRepo.On("FindByID", uint(1)).Return(util.Optional[domain.Greeting]{Value: nil}, nil)

	service := NewHelloService(mockRepo, mockMapper, mockClock)

	_, err := service.GetGreetingByID(1)

	assert.Error(t, err, "An error should be returned when greeting is not found")
	assert.IsType(t, &customError.ResourceNotFoundError{}, err, "Error should be of type ResourceNotFoundError")

	var notFoundErr *customError.ResourceNotFoundError
	if errors.As(err, &notFoundErr) {
		assert.Equal(t, "Greeting", notFoundErr.Resource, "Resource should be 'Greeting'")
		assert.Equal(t, "id", notFoundErr.Criteria, "Criteria should be 'id'")
		assert.Equal(t, "1", notFoundErr.Value, "Value should match the missing ID")
	}

	mockRepo.AssertExpectations(t)
}

func TestHelloService_GetGreetingByID_RepoError(t *testing.T) {
	mockRepo := new(customMock.MockHelloRepository)
	mockMapper := new(customMock.MockHelloMapper)
	mockClock := new(customMock.MockClock)

	expectedError := errors.New("database error")
	mockRepo.On("FindByID", uint(1)).Return(util.Optional[domain.Greeting]{}, expectedError)

	service := NewHelloService(mockRepo, mockMapper, mockClock)

	_, err := service.GetGreetingByID(1)

	assert.Error(t, err, "An error should be returned when repository fails")
	assert.ErrorContains(t, err, "database error", "Error should contain the expected repository error")

	mockRepo.AssertExpectations(t)
}

func TestHelloService_UpdateGreeting_Success(t *testing.T) {
	mockRepo := new(customMock.MockHelloRepository)
	mockMapper := new(customMock.MockHelloMapper)
	mockClock := new(customMock.MockClock)

	// Mock data
	existingEntity := domain.Greeting{ID: 1, Message: "Old Message"}
	updatedEntity := domain.Greeting{ID: 1, Message: "Updated Message"}
	input := dto.GreetingInput{Message: "Updated Message"}
	expectedResponse := dto.GreetingResponse{ID: 1, Message: "Updated Message"}

	// Mock expectations
	mockRepo.On("FindByID", uint(1)).Return(util.Optional[domain.Greeting]{Value: &existingEntity}, nil)
	mockMapper.On("PartialUpdateGreeting", &existingEntity, input)
	mockRepo.On("Save", existingEntity).Return(updatedEntity, nil)
	mockMapper.On("ToGreetingResponse", updatedEntity).Return(expectedResponse, nil)

	service := NewHelloService(mockRepo, mockMapper, mockClock)

	// Call the method under test
	actual, err := service.UpdateGreeting(1, input)

	// Assertions
	assert.NoError(t, err, "There should be no error")
	assert.Equal(t, expectedResponse, actual, "Updated greeting should match the expected response")

	// Verify mock expectations
	mockRepo.AssertExpectations(t)
	mockMapper.AssertExpectations(t)
}

func TestHelloService_UpdateGreeting_NotFound(t *testing.T) {
	mockRepo := new(customMock.MockHelloRepository)
	mockMapper := new(customMock.MockHelloMapper)
	mockClock := new(customMock.MockClock)

	// Mock data
	input := dto.GreetingInput{Message: "Updated Message"}

	// Mock expectations
	mockRepo.On("FindByID", uint(1)).Return(util.Optional[domain.Greeting]{Value: nil}, nil)

	service := NewHelloService(mockRepo, mockMapper, mockClock)

	// Call the method under test
	_, err := service.UpdateGreeting(1, input)

	// Assertions
	assert.Error(t, err, "An error should be returned when greeting is not found")
	assert.IsType(t, &customError.ResourceNotFoundError{}, err, "Error should be of type ResourceNotFoundError")

	var notFoundErr *customError.ResourceNotFoundError
	if errors.As(err, &notFoundErr) {
		assert.Equal(t, "Greeting", notFoundErr.Resource, "Resource should be 'Greeting'")
		assert.Equal(t, "id", notFoundErr.Criteria, "Criteria should be 'id'")
		assert.Equal(t, "1", notFoundErr.Value, "Value should match the missing ID")
	}

	// Verify mock expectations
	mockRepo.AssertExpectations(t)
}

func TestHelloService_UpdateGreeting_RepoError(t *testing.T) {
	mockRepo := new(customMock.MockHelloRepository)
	mockMapper := new(customMock.MockHelloMapper)
	mockClock := new(customMock.MockClock)

	// Mock data
	existingEntity := domain.Greeting{ID: 1, Message: "Old Message"}
	input := dto.GreetingInput{Message: "Updated Message"}

	// Mock expectations
	mockRepo.On("FindByID", uint(1)).Return(util.Optional[domain.Greeting]{Value: &existingEntity}, nil)
	mockMapper.On("PartialUpdateGreeting", &existingEntity, input)
	mockRepo.On("Save", existingEntity).Return(domain.Greeting{}, errors.New("database error"))

	service := NewHelloService(mockRepo, mockMapper, mockClock)

	// Call the method under test
	_, err := service.UpdateGreeting(1, input)

	// Assertions
	assert.Error(t, err, "An error should be returned when repository fails")
	assert.ErrorContains(t, err, "database error", "Error should contain the expected repository error")

	// Verify mock expectations
	mockRepo.AssertExpectations(t)
	mockMapper.AssertExpectations(t)
}

func TestHelloService_DeleteGreeting_Success(t *testing.T) {
	mockRepo := new(customMock.MockHelloRepository)
	mockMapper := new(customMock.MockHelloMapper)
	mockClock := new(customMock.MockClock)

	// Mock data
	existingEntity := domain.Greeting{ID: 1, Message: "Hello, World!"}

	// Mock expectations
	mockRepo.On("FindByID", uint(1)).Return(util.Optional[domain.Greeting]{Value: &existingEntity}, nil)
	mockRepo.On("DeleteByID", uint(1)).Return(nil)

	service := NewHelloService(mockRepo, mockMapper, mockClock)

	// Call the method under test
	err := service.DeleteGreeting(1)

	// Assertions
	assert.NoError(t, err, "There should be no error when deleting a greeting")

	// Verify mock expectations
	mockRepo.AssertExpectations(t)
}

func TestHelloService_DeleteGreeting_NotFound(t *testing.T) {
	mockRepo := new(customMock.MockHelloRepository)
	mockMapper := new(customMock.MockHelloMapper)
	mockClock := new(customMock.MockClock)

	// Mock expectations
	mockRepo.On("FindByID", uint(1)).Return(util.Optional[domain.Greeting]{Value: nil}, nil)

	service := NewHelloService(mockRepo, mockMapper, mockClock)

	// Call the method under test
	err := service.DeleteGreeting(1)

	// Assertions
	assert.Error(t, err, "An error should be returned when greeting is not found")
	assert.IsType(t, &customError.ResourceNotFoundError{}, err, "Error should be of type ResourceNotFoundError")

	var notFoundErr *customError.ResourceNotFoundError
	if errors.As(err, &notFoundErr) {
		assert.Equal(t, "Greeting", notFoundErr.Resource, "Resource should be 'Greeting'")
		assert.Equal(t, "id", notFoundErr.Criteria, "Criteria should be 'id'")
		assert.Equal(t, "1", notFoundErr.Value, "Value should match the missing ID")
	}

	// Verify mock expectations
	mockRepo.AssertExpectations(t)
}

func TestHelloService_DeleteGreeting_RepoError(t *testing.T) {
	mockRepo := new(customMock.MockHelloRepository)
	mockMapper := new(customMock.MockHelloMapper)
	mockClock := new(customMock.MockClock)

	// Mock data
	existingEntity := domain.Greeting{ID: 1, Message: "Hello, World!"}

	// Mock expectations
	mockRepo.On("FindByID", uint(1)).Return(util.Optional[domain.Greeting]{Value: &existingEntity}, nil)
	mockRepo.On("DeleteByID", uint(1)).Return(errors.New("database error"))

	service := NewHelloService(mockRepo, mockMapper, mockClock)

	// Call the method under test
	err := service.DeleteGreeting(1)

	// Assertions
	assert.Error(t, err, "An error should be returned when repository fails to delete")
	assert.ErrorContains(t, err, "database error", "Error should contain the expected repository error")

	// Verify mock expectations
	mockRepo.AssertExpectations(t)
}
