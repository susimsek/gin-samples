package router

import (
	"gin-samples/internal/controller"
	"gin-samples/internal/service"
	"github.com/gin-gonic/gin"
)

func HelloRoutes(r *gin.Engine) {
	helloService := service.NewHelloService()
	helloController := controller.NewHelloController(helloService)

	r.GET("/api/hello", helloController.Hello)
}
