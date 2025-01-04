package service

import (
	customError "gin-samples/internal/error"
	"gin-samples/internal/model"
	"gin-samples/internal/repository"
)

type HelloService interface {
	GetGreeting() model.Greeting
	CreateGreeting(input model.GreetingInput) (model.Greeting, error)
	GetAllGreetings() []model.Greeting
}

type helloServiceImpl struct {
	repo repository.HelloRepository
}

func NewHelloService(repo repository.HelloRepository) HelloService {
	return &helloServiceImpl{
		repo: repo,
	}
}

func (s *helloServiceImpl) GetGreeting() model.Greeting {
	return model.Greeting{Message: "Hello, World!"}
}

func (s *helloServiceImpl) CreateGreeting(input model.GreetingInput) (model.Greeting, error) {
	// Check if a greeting with the same message already exists
	if s.repo.ExistsByMessage(input.Message) {
		return model.Greeting{}, &customError.ResourceConflictError{
			Resource: "Greeting",
			Criteria: "message",
			Value:    input.Message,
		}
	}

	// Save the new greeting
	return s.repo.SaveGreeting(input), nil
}

func (s *helloServiceImpl) GetAllGreetings() []model.Greeting {
	return s.repo.GetAllGreetings()
}
