package mock

import (
	"gin-samples/internal/domain"
	"gin-samples/internal/util"
	"github.com/stretchr/testify/mock"
)

// MockHelloRepository is a mock implementation of HelloRepository
type MockHelloRepository struct {
	mock.Mock
}

// ExistsByMessage checks if a message exists
func (m *MockHelloRepository) ExistsByMessage(message string) (bool, error) {
	args := m.Called(message)
	return args.Bool(0), args.Error(1)
}

// Save saves a greeting
func (m *MockHelloRepository) Save(greeting domain.Greeting) (domain.Greeting, error) {
	args := m.Called(greeting)
	if args.Get(0) == nil {
		return domain.Greeting{}, args.Error(1)
	}
	return args.Get(0).(domain.Greeting), args.Error(1)
}

// FindAll retrieves all greetings
func (m *MockHelloRepository) FindAll() ([]domain.Greeting, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Greeting), args.Error(1)
}

// FindByID retrieves a greeting by its ID and returns an Optional
func (m *MockHelloRepository) FindByID(id uint) (util.Optional[domain.Greeting], error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return util.Optional[domain.Greeting]{Value: nil}, args.Error(1)
	}
	return args.Get(0).(util.Optional[domain.Greeting]), args.Error(1)
}

// DeleteByID deletes a greeting by its ID
func (m *MockHelloRepository) DeleteByID(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
