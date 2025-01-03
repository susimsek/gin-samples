package router

import (
	"gin-samples/internal/controller"

	"github.com/gin-gonic/gin"
)

func AddHelloRoutes(r *gin.Engine, helloController controller.HelloController) {
	r.GET("/api/hello", helloController.Hello)
	r.POST("/api/hello", helloController.CreateGreeting)
}
