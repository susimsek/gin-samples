package mock

import (
	"gin-samples/internal/entity"
	"github.com/stretchr/testify/mock"
)

// MockHelloRepository is a mock implementation of HelloRepository
type MockHelloRepository struct {
	mock.Mock
}

func (m *MockHelloRepository) ExistsByMessage(message string) (bool, error) {
	args := m.Called(message)
	return args.Bool(0), args.Error(1)
}

func (m *MockHelloRepository) SaveGreeting(greeting *entity.Greeting) (*entity.Greeting, error) {
	args := m.Called(greeting)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Greeting), args.Error(1)
}

func (m *MockHelloRepository) GetAllGreetings() ([]entity.Greeting, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Greeting), args.Error(1)
}

func (m *MockHelloRepository) FindByMessage(message string) (*entity.Greeting, error) {
	args := m.Called(message)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Greeting), args.Error(1)
}
