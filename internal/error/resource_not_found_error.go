package error

import "fmt"

// ResourceNotFoundError represents an error for missing resources
type ResourceNotFoundError struct {
	Resource string
	Criteria string
	Value    string
}

func (e *ResourceNotFoundError) Error() string {
	return fmt.Sprintf("The %s could not be found with %s: %s", e.Resource, e.Criteria, e.Value)
}
