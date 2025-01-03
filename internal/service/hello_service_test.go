package service

import (
	"gin-samples/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloService_GetGreeting(t *testing.T) {
	service := NewHelloService()

	expected := model.Greeting{Message: "Hello, World!"}

	actual := service.GetGreeting()

	assert.Equal(t, expected, actual, "Greeting message should match the expected value")
}
