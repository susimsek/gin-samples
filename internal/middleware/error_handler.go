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

// Constants for reusable values
const (
	TypeAboutBlank           = "about:blank"
	ErrorInvalidRequest      = "invalid_request"
	ErrorResourceConflict    = "resource_conflict"
	ErrorInternalServer      = "server_error"
	TitleBadRequest          = "Bad Request"
	TitleConflict            = "Conflict"
	TitleInternalServerError = "Internal Server Error"
)

// ErrorHandlingMiddleware handles and formats errors in the application.
func ErrorHandlingMiddleware(trans ut.Translator) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			problemDetail := handleErrors(c, trans)
			c.JSON(problemDetail.Status, problemDetail)
			c.Abort()
		}
	}
}

func handleErrors(c *gin.Context, trans ut.Translator) model.ProblemDetail {
	for _, err := range c.Errors {
		if problemDetail, ok := handleMessageNotReadableError(err, c); ok {
			return problemDetail
		}
		if problemDetail, ok := handleValidationErrors(err, c, trans); ok {
			return problemDetail
		}
		if problemDetail, ok := handleConflictErrors(err, c); ok {
			return problemDetail
		}
	}

	return handleInternalServerError(c)
}

func handleMessageNotReadableError(err *gin.Error, c *gin.Context) (model.ProblemDetail, bool) {
	var messageNotReadableErr *customError.MessageNotReadableError
	if errors.As(err.Err, &messageNotReadableErr) {
		return model.ProblemDetail{
			Type:     TypeAboutBlank,
			Title:    TitleBadRequest,
			Status:   http.StatusBadRequest,
			Detail:   "The request message could not be read. Please check the format and try again.", // Sabit mesaj
			Error:    ErrorInvalidRequest,
			Instance: c.Request.URL.Path,
		}, true
	}
	return model.ProblemDetail{}, false
}

func handleValidationErrors(err *gin.Error, c *gin.Context, trans ut.Translator) (model.ProblemDetail, bool) {
	var validationErrs validator.ValidationErrors
	if errors.As(err.Err, &validationErrs) {
		var violations []model.Violation
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
		return model.ProblemDetail{
			Type:       TypeAboutBlank,
			Title:      TitleBadRequest,
			Status:     http.StatusBadRequest,
			Detail:     "Validation error occurred.",
			Error:      ErrorInvalidRequest,
			Instance:   c.Request.URL.Path,
			Violations: violations,
		}, true
	}
	return model.ProblemDetail{}, false
}

func handleConflictErrors(err *gin.Error, c *gin.Context) (model.ProblemDetail, bool) {
	var conflictErr *customError.ResourceConflictError
	if errors.As(err.Err, &conflictErr) {
		return model.ProblemDetail{
			Type:     TypeAboutBlank,
			Title:    TitleConflict,
			Status:   http.StatusConflict,
			Detail:   conflictErr.Error(),
			Error:    ErrorResourceConflict,
			Instance: c.Request.URL.Path,
		}, true
	}
	return model.ProblemDetail{}, false
}

func handleInternalServerError(c *gin.Context) model.ProblemDetail {
	return model.ProblemDetail{
		Type:     TypeAboutBlank,
		Title:    TitleInternalServerError,
		Status:   http.StatusInternalServerError,
		Detail:   "An internal server error occurred. Please try again later.",
		Error:    ErrorInternalServer,
		Instance: c.Request.URL.Path,
	}
}
