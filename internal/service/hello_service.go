package service

import (
	"fmt"
	"gin-samples/internal/dto"
	"gin-samples/internal/entity"
	customError "gin-samples/internal/error"
	"gin-samples/internal/repository"
	"gin-samples/internal/utils"
)

type HelloService interface {
	GetGreeting() dto.GreetingResponse
	CreateGreeting(input dto.GreetingInput) (dto.GreetingResponse, error)
	GetAllGreetings() ([]dto.GreetingResponse, error)
}

type helloServiceImpl struct {
	repo  repository.HelloRepository
	clock utils.Clock
}

func NewHelloService(repo repository.HelloRepository,
	clock utils.Clock) HelloService {
	return &helloServiceImpl{
		repo:  repo,
		clock: clock,
	}
}

// GetGreeting returns a static greeting message
func (s *helloServiceImpl) GetGreeting() dto.GreetingResponse {
	now := s.clock.Now()
	return dto.GreetingResponse{
		ID:        0,
		Message:   "Hello, World!",
		CreatedAt: now,
		UpdatedAt: now,
	}
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
		ID:        savedGreeting.ID,
		Message:   savedGreeting.Message,
		CreatedAt: savedGreeting.CreatedAt,
		UpdatedAt: savedGreeting.UpdatedAt,
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
			ID:        greeting.ID,
			Message:   greeting.Message,
			CreatedAt: greeting.CreatedAt,
			UpdatedAt: greeting.UpdatedAt,
		}
	}
	return response, nil
}
