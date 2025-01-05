package mock

import (
	"gin-samples/internal/dto"
	"gin-samples/internal/entity"
	"github.com/stretchr/testify/mock"
)

// MockHelloMapper mocks the HelloMapper interface
type MockHelloMapper struct {
	mock.Mock
}

func (m *MockHelloMapper) ToGreetingResponse(g entity.Greeting) (dto.GreetingResponse, error) {
	args := m.Called(g)
	return args.Get(0).(dto.GreetingResponse), args.Error(1)
}

func (m *MockHelloMapper) ToGreetingResponses(greetings []entity.Greeting) ([]dto.GreetingResponse, error) {
	args := m.Called(greetings)
	return args.Get(0).([]dto.GreetingResponse), args.Error(1)
}

func (m *MockHelloMapper) ToGreetingEntity(input dto.GreetingInput) (entity.Greeting, error) {
	args := m.Called(input)
	return args.Get(0).(entity.Greeting), args.Error(1)
}
