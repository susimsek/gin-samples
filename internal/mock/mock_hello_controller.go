package mock

import (
	"gin-samples/internal/dto"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// MockHelloController simulates the HelloController
type MockHelloController struct{}

// Hello simulates a static greeting response
func (m *MockHelloController) Hello(c *gin.Context) {
	fixedTime := time.Date(2025, 1, 5, 10, 0, 0, 0, time.UTC)
	c.JSON(http.StatusOK, dto.GreetingResponse{
		ID:        1,
		Message:   "Mocked Hello, World!",
		CreatedAt: fixedTime,
		UpdatedAt: fixedTime,
	})
}

// CreateGreeting simulates creating a new greeting
func (m *MockHelloController) CreateGreeting(c *gin.Context) {
	var input dto.GreetingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fixedTime := time.Date(2025, 1, 5, 11, 0, 0, 0, time.UTC)

	// Simulate the creation of a new greeting
	c.JSON(http.StatusCreated, dto.GreetingResponse{
		ID:        2, // Simulated ID
		Message:   input.Message,
		CreatedAt: fixedTime,
		UpdatedAt: fixedTime,
	})
}

// GetAllGreetings simulates retrieving all greetings
func (m *MockHelloController) GetAllGreetings(c *gin.Context) {
	fixedTime := time.Date(2025, 1, 5, 10, 0, 0, 0, time.UTC)

	mockGreetings := []dto.GreetingResponse{
		{ID: 1, Message: "Mocked Hello, World!", CreatedAt: fixedTime, UpdatedAt: fixedTime},
		{ID: 2, Message: "Mocked Hi!", CreatedAt: fixedTime, UpdatedAt: fixedTime},
	}
	c.JSON(http.StatusOK, mockGreetings)
}

// GetGreetingByID simulates retrieving a greeting by its ID
func (m *MockHelloController) GetGreetingByID(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "1" {
		fixedTime := time.Date(2025, 1, 5, 10, 0, 0, 0, time.UTC)
		c.JSON(http.StatusOK, dto.GreetingResponse{
			ID:        1,
			Message:   "Mocked Hello, World!",
			CreatedAt: fixedTime,
			UpdatedAt: fixedTime,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Greeting not found",
		})
	}
}

// UpdateGreeting simulates updating a greeting by its ID
func (m *MockHelloController) UpdateGreeting(c *gin.Context) {
	idParam := c.Param("id")
	if idParam != "1" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Greeting not found",
		})
		return
	}

	var input dto.GreetingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simulate the updated greeting
	fixedTime := time.Date(2025, 1, 5, 11, 30, 0, 0, time.UTC)
	c.JSON(http.StatusOK, dto.GreetingResponse{
		ID:        1,
		Message:   input.Message,
		CreatedAt: time.Date(2025, 1, 5, 10, 0, 0, 0, time.UTC), // Keep original CreatedAt
		UpdatedAt: fixedTime,
	})
}

// DeleteGreeting simulates deleting a greeting by its ID
func (m *MockHelloController) DeleteGreeting(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "1" {
		// Simulate successful deletion
		c.JSON(http.StatusNoContent, gin.H{}) // No Content response
	} else {
		// Simulate not found scenario
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Greeting not found",
		})
	}
}
