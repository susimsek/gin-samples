package error

// AccessDeniedError represents an error for access denial
type AccessDeniedError struct {
	Message string // Error message, must be provided
}

// Error returns the error message
func (e *AccessDeniedError) Error() string {
	return e.Message
}
