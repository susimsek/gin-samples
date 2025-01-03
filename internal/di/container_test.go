package di_test

import (
	"testing"

	"gin-samples/internal/di"
	"github.com/stretchr/testify/assert"
)

func TestNewContainer(t *testing.T) {
	container := di.NewContainer()

	assert.NotNil(t, container)

	assert.NotNil(t, container.HelloService)
	assert.NotNil(t, container.HelloController)
	assert.NotNil(t, container.Router)

	greeting := container.HelloService.GetGreeting()
	assert.Equal(t, "Hello, World!", greeting.Message)
}
