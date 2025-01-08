package error

// JwtError represents an error for invalid or expired JWT token
type JwtError struct {
	Message string // Error message, must be provided
}

// Error returns the error message
func (e *JwtError) Error() string {
	return e.Message
}
