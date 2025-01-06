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
	healthController controller.HealthController,
	trans ut.Translator) *gin.Engine {
	r := gin.Default()
	r.StaticFile("/favicon.ico", "./resources/favicons/favicon.ico")
	r.Use(middleware.ErrorHandlingMiddleware(trans))

	// Add Hello routes
	AddHelloRoutes(r, helloController)

	// Add Health routes
	AddHealthRoutes(r, healthController)

	// Swagger route
	r.GET("/swagger-ui/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
