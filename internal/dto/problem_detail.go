package dto

// Violation represents a single validation error
// @Description Represents a single validation error for a field
type Violation struct {
	// Code of the validation rule that was violated
	Code string `json:"code" example:"min"`

	// Object that contains the field (optional)
	Object string `json:"object,omitempty" example:"GreetingInput"`

	// Field that failed validation
	Field string `json:"field" example:"message"`

	// Rejected value of the field
	RejectedValue string `json:"rejectedValue" example:"Hi"`

	// Error message for the violation
	Message string `json:"message" example:"The message must be at least 3 characters long"`
}

// ProblemDetail represents the structure of error responses
// @Description Represents a structured error response for the API
type ProblemDetail struct {
	// Type of the error, usually a URI identifying the error type
	Type string `json:"type" example:"about:blank"`

	// Title is a short, human-readable summary of the error
	Title string `json:"title" example:"Validation Error"`

	// HTTP status code for the error
	Status int `json:"status" example:"400"`

	// Detail provides a more detailed explanation of the error
	Detail string `json:"detail" example:"One or more fields failed validation"`

	// Instance is a URI that identifies the specific occurrence of the error
	Instance string `json:"instance" example:"/api/hello"`

	// Error is a machine-readable error code
	Error string `json:"error" example:"invalid_request"`

	// Violations is a list of validation errors (optional)
	Violations []Violation `json:"violations,omitempty"`
}
