package middleware

import (
	customError "gin-samples/internal/error"
	"gin-samples/internal/security"
	"github.com/gin-gonic/gin"
)

// AuthorityMiddleware checks if the user has the required authority (e.g., "ROLE_ADMIN")
func AuthorityMiddleware(requiredAuthority string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the claims from the context, which were set in AuthMiddleware
		claims, exists := c.Get("jwt")
		if !exists {
			// If there are no JWT claims in the context, return JwtError
			_ = c.Error(&customError.JwtError{Message: "Invalid or missing JWT token"})
			c.Abort()
			return
		}

		// Assuming claims is of type TokenClaims
		tokenClaims, ok := claims.(*security.TokenClaims)
		if !ok {
			// If the claims are not of type TokenClaims, return JwtError
			_ = c.Error(&customError.JwtError{Message: "Invalid JWT claims"})
			c.Abort()
			return
		}

		// Extract authorities from the token claims
		authorities := tokenClaims.Authorities
		if len(authorities) == 0 {
			// If "authorities" claim is missing or empty, return AccessDeniedError
			_ = c.Error(&customError.AccessDeniedError{Message: "Access Denied: Missing authorities claim"})
			c.Abort()
			return
		}

		// Check if the required authority exists in the authorities
		hasAuthority := false
		for _, authority := range authorities {
			if authority == requiredAuthority {
				hasAuthority = true
				break
			}
		}

		// If the user does not have the required authority, return AccessDeniedError
		if !hasAuthority {
			_ = c.Error(&customError.AccessDeniedError{Message: "Access Denied: Insufficient permissions"})
			c.Abort()
			return
		}

		// Continue to the next handler if the user has the correct authority
		c.Next()
	}
}
