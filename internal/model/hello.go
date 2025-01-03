package model

// Greeting represents a greeting message
// @Description Greeting model
type Greeting struct {
	// Message is the greeting text
	Message string `json:"message" example:"Hello, World!" minLength:"1" maxLength:"100" validate:"required"`
}
