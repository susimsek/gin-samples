package di

import (
	"gin-samples/internal/controller"
	"gin-samples/internal/router"
	"gin-samples/internal/service"

	"github.com/gin-gonic/gin"
)

type Container struct {
	HelloService    service.HelloService
	HelloController controller.HelloController
	Router          *gin.Engine
}

func NewContainer() *Container {
	// Service
	helloService := service.NewHelloService()

	// Controller
	helloController := controller.NewHelloController(helloService)

	// Router
	r := router.SetupRouter(helloController)

	return &Container{
		HelloService:    helloService,
		HelloController: helloController,
		Router:          r,
	}
}
