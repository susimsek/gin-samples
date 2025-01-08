package router

import (
	"gin-samples/internal/controller"

	"github.com/gin-gonic/gin"
)

// AddAuthRoutes adds authentication routes to the router
func AddAuthRoutes(r *gin.Engine, authController controller.AuthenticationController) {
	r.POST("/api/auth/login", authController.Login)
}
