package dto

import "time"

// GreetingResponse represents a greeting message
// @Description Greeting dto
type GreetingResponse struct {
	// ID of the greeting
	ID uint `json:"id" example:"1" validate:"required"`

	// Message is the greeting text
	Message string `json:"message" example:"Hello, World!" minLength:"3" maxLength:"100" validate:"required"`

	// CreatedAt is the timestamp when the greeting was created
	CreatedAt time.Time `json:"createdAt" example:"2025-01-05T10:00:00Z" validate:"required"`

	// UpdatedAt is the timestamp when the greeting was last updated
	UpdatedAt time.Time `json:"updatedAt" example:"2025-01-05T12:00:00Z"`
}

// GreetingInput represents the input for creating a greeting
// @Description Input dto for creating a new greeting
type GreetingInput struct {
	// Message is the greeting text to be created
	Message string `json:"message" example:"Hello, World!" minLength:"3" maxLength:"100" validate:"required,min=3,max=100"`
}
