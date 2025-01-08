package middleware

import (
	customError "gin-samples/internal/error"
	"gin-samples/internal/security"
	"github.com/gin-gonic/gin"
	"regexp"
)

// Regex pattern for Authorization header
var authorizationPattern = regexp.MustCompile(`^Bearer (?P<token>[a-zA-Z0-9-._~+/]+=*)$`)

// AuthMiddleware validates the JWT token from the Authorization header and sets claims in the context.
func AuthMiddleware(tokenGenerator security.TokenGenerator) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			_ = c.Error(&customError.JwtError{Message: "Missing Authorization header"})
			c.Abort()
			return
		}

		// Match the pattern
		matches := authorizationPattern.FindStringSubmatch(authHeader)
		if len(matches) == 0 {
			_ = c.Error(&customError.JwtError{Message: "Invalid Authorization header format"})
			c.Abort()
			return
		}

		// Extract the token from the match
		tokenIndex := authorizationPattern.SubexpIndex("token")
		if tokenIndex == -1 || len(matches) <= tokenIndex {
			_ = c.Error(&customError.JwtError{Message: "Token not found in Authorization header"})
			c.Abort()
			return
		}
		token := matches[tokenIndex]

		// Validate the token
		claims, err := tokenGenerator.Validate(token)
		if err != nil {
			_ = c.Error(&customError.JwtError{Message: "Invalid or expired token"})
			c.Abort()
			return
		}

		// Add the entire claims to the context
		c.Set("jwt", claims)

		// If the token is valid, continue to the next handler
		c.Next()
	}
}
