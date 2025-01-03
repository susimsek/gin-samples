package mock

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MockHelloController struct{}

func (m *MockHelloController) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Mocked Hello, World!"})
}
