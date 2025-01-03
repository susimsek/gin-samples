package mock

import (
	"gin-samples/internal/di"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewMockContainer() *di.Container {
	r := gin.Default()
	r.GET("/mocked", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "This is a mocked container"})
	})

	return &di.Container{
		HelloService:    nil,
		HelloController: nil,
		Router:          r,
	}
}
