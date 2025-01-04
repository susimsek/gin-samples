package model

// Violation represents a single validation error
type Violation struct {
	Code          string `json:"code"`
	Object        string `json:"object"`
	Field         string `json:"field"`
	RejectedValue string `json:"rejectedValue"`
	Message       string `json:"message"`
}

// ProblemDetail represents the structure of error responses
type ProblemDetail struct {
	Type       string      `json:"type"`
	Title      string      `json:"title"`
	Status     int         `json:"status"`
	Detail     string      `json:"detail"`
	Instance   string      `json:"instance"`
	Error      string      `json:"error"`
	Violations []Violation `json:"violations,omitempty"`
}
