package dto

// GreetingResponse represents a greeting message
// @Description Greeting dto
type GreetingResponse struct {
	// ID of the greeting
	ID uint `json:"id" example:"1" validate:"required"`
	// Message is the greeting text
	Message string `json:"message" example:"Hello, World!" minLength:"3" maxLength:"100" validate:"required"`
}

// GreetingInput represents the input for creating a greeting
// @Description Input dto for creating a new greeting
type GreetingInput struct {
	// Message is the greeting text to be created
	Message string `json:"message" example:"Hello, World!" minLength:"3" maxLength:"100" validate:"required,min=3,max=100"`
}
