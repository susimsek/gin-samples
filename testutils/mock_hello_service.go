package testutils

import "gin-samples/internal/model"

type MockHelloService struct{}

func (m *MockHelloService) GetGreeting() model.Greeting {
	return model.Greeting{Message: "Mock Hello"}
}
