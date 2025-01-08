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
	// Group for authenticated users (all users who have a valid JWT)
	authenticatedGroup := r.Group("/api")
	authenticatedGroup.Use(middleware.AuthMiddleware(tokenGenerator))

	// Create an admin-specific group with additional access controls (admin check)
	adminGroup := r.Group("/api")
	adminGroup.Use(middleware.AuthMiddleware(tokenGenerator))
	adminGroup.Use(middleware.AuthorityMiddleware("ROLE_ADMIN")) // Ensures only admin has access to this group

	// Add Hello routes
	AddHelloRoutes(authenticatedGroup, helloController)

	// Add Health routes
	AddHealthRoutes(r, healthController)

	// Add Authentication routes
	AddAuthRoutes(r, authController)

	AddAdminRoutes(adminGroup, helloController)

	// Swagger route
	r.GET("/swagger-ui/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
