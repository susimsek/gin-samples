package mock

import (
	"github.com/stretchr/testify/mock"
	"time"
)

// MockClock is a mock implementation of the Clock interface
type MockClock struct {
	mock.Mock
}

// Now returns a mocked current time
func (m *MockClock) Now() time.Time {
	args := m.Called()
	return args.Get(0).(time.Time)
}
