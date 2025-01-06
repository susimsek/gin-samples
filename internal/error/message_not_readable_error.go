package error

import "fmt"

// MessageNotReadableError represents an error when a request body cannot be read
type MessageNotReadableError struct {
	Detail string
}

func (e *MessageNotReadableError) Error() string {
	return fmt.Sprintf("The request message could not be read. Detail: %s", e.Detail)
}
