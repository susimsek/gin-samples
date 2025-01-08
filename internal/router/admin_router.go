package router

import (
	"gin-samples/internal/controller"
	"github.com/gin-gonic/gin"
)

// AddAdminRoutes sets up Admin-specific API routes
func AddAdminRoutes(r *gin.RouterGroup, helloController controller.HelloController) {
	// Admin-only route for /hello in the admin group
	r.GET("/hello", helloController.Hello) // Admin users only (adminGroup)
	// You can add more admin-specific routes here
}
