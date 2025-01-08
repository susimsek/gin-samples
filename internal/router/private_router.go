package router

import (
	"gin-samples/internal/middleware"
	"gin-samples/internal/security"
	"github.com/gin-gonic/gin"
)

// SetupPrivateRoutes sets up authenticated routes with AuthMiddleware, excluding specified paths
func SetupPrivateRoutes(r *gin.Engine, tokenGenerator security.TokenGenerator) (*gin.RouterGroup, *gin.RouterGroup) {
	// Group for authenticated users (all users who have a valid JWT)
	authenticatedGroup := r.Group("/api")
	authenticatedGroup.Use(middleware.AuthMiddleware(tokenGenerator))

	// Create an admin-specific group with additional access controls (admin check)
	adminGroup := r.Group("/api")
	adminGroup.Use(middleware.AuthMiddleware(tokenGenerator))
	adminGroup.Use(middleware.AuthorityMiddleware("ROLE_ADMIN")) // Ensures only admin has access to this group

	return authenticatedGroup, adminGroup
}
