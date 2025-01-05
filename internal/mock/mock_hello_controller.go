package mock

import (
	"gin-samples/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MockHelloController simulates the HelloController
type MockHelloController struct{}

// Hello simulates a static greeting response
func (m *MockHelloController) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, dto.GreetingResponse{
		ID:      1,
		Message: "Mocked Hello, World!",
	})
}

// CreateGreeting simulates creating a new greeting
func (m *MockHelloController) CreateGreeting(c *gin.Context) {
	var input dto.GreetingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simulate the creation of a new greeting
	c.JSON(http.StatusCreated, dto.GreetingResponse{
		ID:      2, // Simulated ID
		Message: input.Message,
	})
}

// GetAllGreetings simulates retrieving all greetings
func (m *MockHelloController) GetAllGreetings(c *gin.Context) {
	mockGreetings := []dto.GreetingResponse{
		{ID: 1, Message: "Mocked Hello, World!"},
		{ID: 2, Message: "Mocked Hi!"},
	}
	c.JSON(http.StatusOK, mockGreetings)
}
