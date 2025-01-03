package router

import (
	"gin-samples/internal/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(helloController controller.HelloController) *gin.Engine {
	r := gin.Default()

	AddHelloRoutes(r, helloController)

	return r
}
