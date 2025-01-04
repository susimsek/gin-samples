package router

import (
	"gin-samples/internal/controller"

	"github.com/gin-gonic/gin"
)

func AddHealthRoutes(r *gin.Engine, healthController controller.HealthController) {
	r.GET("/health/liveness", healthController.Liveness)
	r.GET("/health/readiness", healthController.Readiness)
}
