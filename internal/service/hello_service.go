package service

import "gin-samples/internal/model"

type HelloService interface {
	GetGreeting() model.Greeting
}

type helloServiceImpl struct{}

func NewHelloService() HelloService {
	return &helloServiceImpl{}
}

func (s *helloServiceImpl) GetGreeting() model.Greeting {
	return model.Greeting{Message: "Hello, World!"}
}
