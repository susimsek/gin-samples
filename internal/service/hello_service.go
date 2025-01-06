package service

import (
	"fmt"
	"gin-samples/internal/dto"
	customError "gin-samples/internal/error"
	"gin-samples/internal/mapper"
	"gin-samples/internal/repository"
	"gin-samples/internal/utils"
)

type HelloService interface {
	GetGreeting() dto.GreetingResponse
	CreateGreeting(input dto.GreetingInput) (dto.GreetingResponse, error)
	GetAllGreetings() ([]dto.GreetingResponse, error)
}

type helloServiceImpl struct {
	repo   repository.HelloRepository
	clock  utils.Clock
	mapper mapper.HelloMapper
}

// NewHelloService creates a new instance of helloServiceImpl
func NewHelloService(repo repository.HelloRepository,
	mapper mapper.HelloMapper,
	clock utils.Clock) HelloService {
	return &helloServiceImpl{
		repo:   repo,
		clock:  clock,
		mapper: mapper,
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

	// Map input DTO to domain
	entity, err := s.mapper.ToGreetingEntity(input)
	if err != nil {
		return dto.GreetingResponse{}, fmt.Errorf("failed to map input: %w", err)
	}

	// Save domain
	savedEntity, err := s.repo.SaveGreeting(entity)
	if err != nil {
		return dto.GreetingResponse{}, fmt.Errorf("failed to save greeting: %w", err)
	}

	// Map saved domain to response DTO
	response, err := s.mapper.ToGreetingResponse(savedEntity)
	if err != nil {
		return dto.GreetingResponse{}, fmt.Errorf("failed to map domain: %w", err)
	}

	return response, nil
}

// GetAllGreetings retrieves all greetings
func (s *helloServiceImpl) GetAllGreetings() ([]dto.GreetingResponse, error) {
	entities, err := s.repo.GetAllGreetings()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch greetings: %w", err)
	}

	// Map entities to response DTOs
	responses, err := s.mapper.ToGreetingResponses(entities)
	if err != nil {
		return nil, fmt.Errorf("failed to map entities: %w", err)
	}

	return responses, nil
}
