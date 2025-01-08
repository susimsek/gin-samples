package mock

import (
	"gin-samples/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MockAuthenticationController is a mock implementation of the AuthenticationController interface
type MockAuthenticationController struct{}

// Login is a mock implementation for login endpoint
func (m *MockAuthenticationController) Login(c *gin.Context) {
	var input dto.LoginInput

	// Mock request binding
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	// Mock behavior
	if input.Username == "admin" && input.Password == "password" {
		c.JSON(http.StatusOK, dto.TokenResponse{
			AccessToken:          "mocked-jwt-token",
			TokenType:            "Bearer",
			AccessTokenExpiresIn: 3600,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password"})
	}
}
