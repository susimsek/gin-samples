package mock

import (
	"gin-samples/internal/domain"
	"gin-samples/internal/dto"
	"github.com/stretchr/testify/mock"
)

// MockHelloMapper mocks the HelloMapper interface
type MockHelloMapper struct {
	mock.Mock
}

func (m *MockHelloMapper) ToGreetingResponse(g domain.Greeting) dto.GreetingResponse {
	args := m.Called(g)
	return args.Get(0).(dto.GreetingResponse)
}

func (m *MockHelloMapper) ToGreetingResponses(greetings []domain.Greeting) []dto.GreetingResponse {
	args := m.Called(greetings)
	return args.Get(0).([]dto.GreetingResponse)
}

func (m *MockHelloMapper) ToGreetingEntity(input dto.GreetingInput) domain.Greeting {
	args := m.Called(input)
	return args.Get(0).(domain.Greeting)
}

func (m *MockHelloMapper) PartialUpdateGreeting(entity *domain.Greeting, input dto.GreetingInput) {
	m.Called(entity, input)
}
