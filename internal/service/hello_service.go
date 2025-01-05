package service

import (
	"fmt"
	"gin-samples/internal/dto"
	"gin-samples/internal/entity"
	customError "gin-samples/internal/error"
	"gin-samples/internal/repository"
)

type HelloService interface {
	GetGreeting() dto.GreetingResponse
	CreateGreeting(input dto.GreetingInput) (dto.GreetingResponse, error)
	GetAllGreetings() ([]dto.GreetingResponse, error)
}

type helloServiceImpl struct {
	repo repository.HelloRepository
}

func NewHelloService(repo repository.HelloRepository) HelloService {
	return &helloServiceImpl{
		repo: repo,
	}
}

// GetGreeting returns a static greeting message
func (s *helloServiceImpl) GetGreeting() dto.GreetingResponse {
	return dto.GreetingResponse{ID: 0, Message: "Hello, World!"}
}

// CreateGreeting creates a new greeting
func (s *helloServiceImpl) CreateGreeting(input dto.GreetingInput) (dto.GreetingResponse, error) {
	// Check if a greeting with the same message already exists
	exists, err := s.repo.ExistsByMessage(input.Message)
	if err != nil {
		return dto.GreetingResponse{}, fmt.Errorf("failed to check existence: %w", err)
	}
	if exists {
		return dto.GreetingResponse{}, &customError.ResourceConflictError{
			Resource: "Greeting",
			Criteria: "message",
			Value:    input.Message,
		}
	}

	// Save the new greeting
	greeting := &entity.Greeting{
		Message: input.Message,
	}
	savedGreeting, err := s.repo.SaveGreeting(greeting)
	if err != nil {
		return dto.GreetingResponse{}, fmt.Errorf("failed to save greeting: %w", err)
	}

	// Convert entity to DTO
	return dto.GreetingResponse{
		ID:      savedGreeting.ID,
		Message: savedGreeting.Message,
	}, nil
}

// GetAllGreetings retrieves all greetings
func (s *helloServiceImpl) GetAllGreetings() ([]dto.GreetingResponse, error) {
	greetings, err := s.repo.GetAllGreetings()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch greetings: %w", err)
	}

	// Convert entities to DTOs
	response := make([]dto.GreetingResponse, len(greetings))
	for i, greeting := range greetings {
		response[i] = dto.GreetingResponse{
			ID:      greeting.ID,
			Message: greeting.Message,
		}
	}
	return response, nil
}
