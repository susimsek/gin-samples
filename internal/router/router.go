package router

import (
	"gin-samples/internal/controller"
	"gin-samples/internal/middleware"
	"gin-samples/internal/security"
	ut "github.com/go-playground/universal-translator"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func SetupRouter(helloController controller.HelloController,
	healthController controller.HealthController,
	authController controller.AuthenticationController,
	trans ut.Translator,
	tokenGenerator security.TokenGenerator) *gin.Engine {
	r := gin.Default()
	r.StaticFile("/favicon.ico", "./resources/favicons/favicon.ico")
	r.Use(middleware.ErrorHandlingMiddleware(trans))
	privateGroup := SetupPrivateRoutes(r, tokenGenerator)
	// Add Hello routes
	AddHelloRoutes(privateGroup, helloController)

	// Add Health routes
	AddHealthRoutes(r, healthController)

	// Add Authentication routes
	AddAuthRoutes(r, authController)

	// Swagger route
	r.GET("/swagger-ui/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
