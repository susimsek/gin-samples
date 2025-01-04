package model

// Greeting represents a greeting message
// @Description Greeting model
type Greeting struct {
	// Message is the greeting text
	Message string `json:"message" example:"Hello, World!" minLength:"1" maxLength:"100" validate:"required"`
}

// GreetingInput represents the input for creating a greeting
// @Description Input model for creating a new greeting
type GreetingInput struct {
	// Message is the greeting text to be created
	Message string `json:"message" example:"Hello, World!" minLength:"1" maxLength:"100" validate:"required,min=3,max=100"`
}
