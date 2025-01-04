package middleware

import (
	"errors"
	customError "gin-samples/internal/error"
	"gin-samples/internal/model"
	ut "github.com/go-playground/universal-translator"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandlingMiddleware(trans ut.Translator) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			var problemDetail model.ProblemDetail
			var violations []model.Violation

			for _, err := range c.Errors {
				var validationErrs validator.ValidationErrors
				if errors.As(err.Err, &validationErrs) {
					for _, ve := range validationErrs {
						translatedMessage := ve.Translate(trans)
						violations = append(violations, model.Violation{
							Code:          ve.Tag(),
							Object:        ve.StructNamespace(),
							Field:         ve.Field(),
							RejectedValue: ve.Param(),
							Message:       translatedMessage,
						})
					}
					problemDetail = model.ProblemDetail{
						Type:       "about:blank",
						Title:      "Bad Request",
						Status:     http.StatusBadRequest,
						Detail:     "Validation error occurred.",
						Error:      "invalid_request",
						Instance:   c.Request.URL.Path,
						Violations: violations,
					}
					break
				}

				var conflictErr *customError.ResourceConflictError
				if errors.As(err.Err, &conflictErr) {
					problemDetail = model.ProblemDetail{
						Type:     "about:blank",
						Title:    "Conflict",
						Status:   http.StatusConflict,
						Detail:   conflictErr.Error(),
						Error:    "resource_conflict",
						Instance: c.Request.URL.Path,
					}
					break
				}

				problemDetail = model.ProblemDetail{
					Type:     "about:blank",
					Title:    "Internal Server Error",
					Status:   http.StatusInternalServerError,
					Detail:   "An internal server error occurred. Please try again later.",
					Error:    "server_error",
					Instance: c.Request.URL.Path,
				}
			}

			c.JSON(problemDetail.Status, problemDetail)
			c.Abort()
		}
	}
}
