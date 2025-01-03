package router

import (
	"gin-samples/internal/controller"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func SetupRouter(helloController controller.HelloController) *gin.Engine {
	r := gin.Default()

	AddHelloRoutes(r, helloController)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
