package mock

import (
	"gin-samples/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MockHelloController struct{}

func (m *MockHelloController) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Mocked Hello, World!"})
}

func (m *MockHelloController) CreateGreeting(c *gin.Context) {
	var input model.GreetingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, model.Greeting{Message: input.Message})
}
