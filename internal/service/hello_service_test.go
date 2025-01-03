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

func TestHelloService_CreateGreeting(t *testing.T) {
	service := NewHelloService()

	input := model.GreetingInput{Message: "Test Greeting"}
	expected := model.Greeting{Message: "Test Greeting"}

	actual := service.CreateGreeting(input)

	assert.Equal(t, expected, actual, "Created greeting should match the input message")
}
