package service

import (
	"gin-samples/internal/model"
)

type HelloService interface {
	GetGreeting() model.Greeting
	CreateGreeting(input model.GreetingInput) model.Greeting
}

type helloServiceImpl struct{}

func NewHelloService() HelloService {
	return &helloServiceImpl{}
}

func (s *helloServiceImpl) GetGreeting() model.Greeting {
	return model.Greeting{Message: "Hello, World!"}
}

func (s *helloServiceImpl) CreateGreeting(input model.GreetingInput) model.Greeting {
	newGreeting := model.Greeting{
		Message: input.Message,
	}
	return newGreeting
}
