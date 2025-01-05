package di_test

import (
	"gin-samples/config"
	"gin-samples/internal/mock"
	"testing"

	"gin-samples/internal/di"
	"github.com/stretchr/testify/assert"
)

func TestNewContainer(t *testing.T) {
	// Create a new container
	config.DatabaseConfig = &mock.MockDatabaseConfig{}
	container := di.NewContainer()

	// Ensure the container is not nil
	assert.NotNil(t, container, "Container should not be nil")

	// Verify Validator and Translator
	assert.NotNil(t, container.Validator, "Validator should not be nil")
	assert.NotNil(t, container.Translator, "Translator should not be nil")

	// Check HelloRepository
	assert.NotNil(t, container.HelloRepository, "HelloRepository should not be nil")

	// Check HelloService
	assert.NotNil(t, container.HelloService, "HelloService should not be nil")

	// Check HelloController
	assert.NotNil(t, container.HelloController, "HelloController should not be nil")

	// Check HealthController
	assert.NotNil(t, container.HealthController, "HealthController should not be nil")

	// Check Router
	assert.NotNil(t, container.Router, "Router should not be nil")
}
