package service

import (
	"fmt"
	"gin-samples/internal/dto"
	customError "gin-samples/internal/error"
	"gin-samples/internal/mapper"
	"gin-samples/internal/repository"
	"gin-samples/internal/util"
)

type HelloService interface {
	GetGreeting() dto.GreetingResponse
	CreateGreeting(input dto.GreetingInput) (dto.GreetingResponse, error)
	GetAllGreetings() ([]dto.GreetingResponse, error)
	GetGreetingByID(id uint) (dto.GreetingResponse, error)
	UpdateGreeting(id uint, input dto.GreetingInput) (dto.GreetingResponse, error)
}

type helloServiceImpl struct {
	repo   repository.HelloRepository
	clock  util.Clock
	mapper mapper.HelloMapper
}

// NewHelloService creates a new instance of helloServiceImpl
func NewHelloService(repo repository.HelloRepository,
	mapper mapper.HelloMapper,
	clock util.Clock) HelloService {
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

	entity := s.mapper.ToGreetingEntity(input)
	savedEntity, err := s.repo.Save(entity)
	if err != nil {
		return dto.GreetingResponse{}, fmt.Errorf("failed to save greeting: %w", err)
	}

	return s.mapper.ToGreetingResponse(savedEntity), nil
}

// GetAllGreetings retrieves all greetings
func (s *helloServiceImpl) GetAllGreetings() ([]dto.GreetingResponse, error) {
	entities, err := s.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch greetings: %w", err)
	}

	return s.mapper.ToGreetingResponses(entities), nil
}

// GetGreetingByID retrieves a greeting by its ID
func (s *helloServiceImpl) GetGreetingByID(id uint) (dto.GreetingResponse, error) {
	optionalEntity, err := s.repo.FindByID(id)
	if err != nil {
		return dto.GreetingResponse{}, fmt.Errorf("failed to fetch greeting by ID: %w", err)
	}

	if optionalEntity.IsEmpty() {
		return dto.GreetingResponse{}, &customError.ResourceNotFoundError{
			Resource: "Greeting",
			Criteria: "id",
			Value:    fmt.Sprintf("%d", id),
		}
	}

	return s.mapper.ToGreetingResponse(*optionalEntity.Value), nil
}

// UpdateGreeting updates an existing greeting by ID
func (s *helloServiceImpl) UpdateGreeting(id uint, input dto.GreetingInput) (dto.GreetingResponse, error) {
	// Fetch the existing greeting
	optionalEntity, err := s.repo.FindByID(id)
	if err != nil {
		return dto.GreetingResponse{}, fmt.Errorf("failed to fetch greeting by ID: %w", err)
	}

	if optionalEntity.IsEmpty() {
		return dto.GreetingResponse{}, &customError.ResourceNotFoundError{
			Resource: "Greeting",
			Criteria: "id",
			Value:    fmt.Sprintf("%d", id),
		}
	}

	// Apply partial update using the mapper
	existingEntity := optionalEntity.Value
	s.mapper.PartialUpdateGreeting(existingEntity, input)

	// Save the updated entity
	updatedEntity, err := s.repo.Save(*existingEntity)
	if err != nil {
		return dto.GreetingResponse{}, fmt.Errorf("failed to update greeting: %w", err)
	}

	// Map the updated entity to response DTO
	return s.mapper.ToGreetingResponse(updatedEntity), nil
}
