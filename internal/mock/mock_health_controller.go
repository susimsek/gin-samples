package mock

import (
	"gin-samples/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MockHealthController struct{}

func (m *MockHealthController) Liveness(c *gin.Context) {
	c.JSON(http.StatusOK, model.HealthStatus{
		Status: "UP",
	})
}

func (m *MockHealthController) Readiness(c *gin.Context) {
	c.JSON(http.StatusOK, model.HealthStatus{
		Status: "UP",
	})
}
