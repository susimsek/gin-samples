package error

import (
	"fmt"
	"gin-samples/internal/dto"
	"strings"
)

// ConstraintViolationError represents a collection of validation violations
type ConstraintViolationError struct {
	Violations []dto.Violation
}

// Error returns a detailed error message including all violations
func (e ConstraintViolationError) Error() string {
	if len(e.Violations) == 0 {
		return "Constraint violation occurred"
	}

	var violationsDetails []string
	for _, v := range e.Violations {
		violationDetail := fmt.Sprintf("Field '%s' (value: '%v') failed validation with rule '%s': %s",
			v.Field, v.RejectedValue, v.Code, v.Message)
		violationsDetails = append(violationsDetails, violationDetail)
	}

	return fmt.Sprintf("Constraint violation occurred: \n%s", strings.Join(violationsDetails, "\n"))
}
