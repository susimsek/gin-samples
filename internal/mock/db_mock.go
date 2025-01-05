package mock

import (
	"gorm.io/gorm"
)

// MockDatabaseConfig is a mock implementation of DatabaseInitializer
type MockDatabaseConfig struct{}

func (m *MockDatabaseConfig) InitDB() *gorm.DB {
	return &gorm.DB{} // Return a fake or nil DB instance
}
