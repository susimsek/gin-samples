package controller

import (
	"gin-samples/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthController interface {
	Liveness(c *gin.Context)
	Readiness(c *gin.Context)
}

type healthControllerImpl struct{}

func NewHealthController() HealthController {
	return &healthControllerImpl{}
}

// Liveness godoc
// @Summary Check if the application is alive
// @Description Returns the liveness status of the application
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} model.HealthStatus
// @Router /health/liveness [get]
func (h *healthControllerImpl) Liveness(c *gin.Context) {
	c.JSON(http.StatusOK, model.HealthStatus{
		Status: "UP",
	})
}

// Readiness godoc
// @Summary Check if the application is ready
// @Description Returns the readiness status of the application
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} model.HealthStatus
// @Router /health/readiness [get]
func (h *healthControllerImpl) Readiness(c *gin.Context) {
	c.JSON(http.StatusOK, model.HealthStatus{
		Status: "UP",
	})
}
