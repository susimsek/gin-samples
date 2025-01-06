package router

import (
	"gin-samples/internal/controller"

	"github.com/gin-gonic/gin"
)

func AddHelloRoutes(r *gin.Engine, helloController controller.HelloController) {
	r.GET("/api/hello", helloController.Hello)
	r.GET("/api/hello/:id", helloController.GetGreetingByID)
	r.POST("/api/hello", helloController.CreateGreeting)
	r.GET("/api/hello/all", helloController.GetAllGreetings)
	r.PUT("/api/hello/:id", helloController.UpdateGreeting)
	r.DELETE("/api/hello/:id", helloController.DeleteGreeting)
}
