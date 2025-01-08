package error

// InvalidCredentialsError represents an error for invalid username or password
type InvalidCredentialsError struct{}

func (e *InvalidCredentialsError) Error() string {
	return "Invalid credentials provided"
}
