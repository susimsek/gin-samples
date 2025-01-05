package mock

import (
	"gin-samples/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MockHealthController struct{}

func (m *MockHealthController) Liveness(c *gin.Context) {
	c.JSON(http.StatusOK, dto.HealthStatus{
		Status: "UP",
	})
}

func (m *MockHealthController) Readiness(c *gin.Context) {
	c.JSON(http.StatusOK, dto.HealthStatus{
		Status: "UP",
	})
}
