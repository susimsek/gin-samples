package router

import (
	"gin-samples/internal/middleware"
	"gin-samples/internal/security"

	"github.com/gin-gonic/gin"
)

// SetupPrivateRoutes sets up private routes with AuthMiddleware, excluding specified paths
func SetupPrivateRoutes(r *gin.Engine, tokenGenerator security.TokenGenerator) *gin.RouterGroup {
	privateGroup := r.Group("/api")
	privateGroup.Use(middleware.AuthMiddleware(tokenGenerator))
	return privateGroup
}
