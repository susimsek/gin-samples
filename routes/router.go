package routes

import (
	"gin-samples/internal/controller"
	"gin-samples/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Dependency Injection
	helloService := service.NewHelloService()
	helloController := controller.NewHelloController(helloService)

	r.GET("/api/hello", helloController.Hello)
	return r
}
