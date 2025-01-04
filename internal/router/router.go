package router

import (
	"gin-samples/internal/controller"
	"gin-samples/internal/middleware"
	ut "github.com/go-playground/universal-translator"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func SetupRouter(helloController controller.HelloController,
	trans ut.Translator) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.ErrorHandlingMiddleware(trans))
	AddHelloRoutes(r, helloController)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
