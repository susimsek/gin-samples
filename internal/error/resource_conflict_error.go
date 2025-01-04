package errors

import "fmt"

// ResourceConflictError represents a conflict error for existing resources
type ResourceConflictError struct {
	Resource string
	Criteria string
	Value    string
}

func (e *ResourceConflictError) Error() string {
	return fmt.Sprintf("%s already exists with %s: %s", e.Resource, e.Criteria, e.Value)
}
