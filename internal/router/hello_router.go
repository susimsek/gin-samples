package router

import (
	"gin-samples/internal/controller"
	"gin-samples/internal/middleware"

	"github.com/gin-gonic/gin"
)

// AddHelloRoutes sets up Hello API routes
func AddHelloRoutes(r *gin.RouterGroup, helloController controller.HelloController) {
	r.GET("/hello", middleware.AuthorityMiddleware("ROLE_ADMIN"),
		helloController.Hello) // Get a greeting message
	r.GET("/hello/:id", helloController.GetGreetingByID)   // Get a greeting by ID
	r.POST("/hello", helloController.CreateGreeting)       // Create a new greeting
	r.GET("/hello/all", helloController.GetAllGreetings)   // Get all greetings
	r.PUT("/hello/:id", helloController.UpdateGreeting)    // Update a greeting by ID
	r.DELETE("/hello/:id", helloController.DeleteGreeting) // Delete a greeting by ID
}
