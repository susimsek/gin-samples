package testutils

import (
	"gin-samples/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MockHelloController struct{}

func (m *MockHelloController) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Mocked Hello, World!"})
}

type MockHelloService struct{}

func (m *MockHelloService) GetGreeting() model.Greeting {
	return model.Greeting{Message: "Mock Hello"}
}
