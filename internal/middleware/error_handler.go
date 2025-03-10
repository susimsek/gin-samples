package middleware

import (
	"errors"
	"gin-samples/internal/dto"
	customError "gin-samples/internal/error"
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
	ErrorResourceNotFound    = "resource_not_found"
	ErrorAccessDenied        = "access_denied"
	ErrorInternalServer      = "server_error"
	TitleBadRequest          = "Bad Request"
	TitleUnauthorized        = "Unauthorized"
	TitleAccessDenied        = "Access Denied"
	TitleConflict            = "Conflict"
	TitleNotFound            = "Not Found"
	TitleInternalServerError = "Internal Server Error"
	DetailValidationError    = "Validation error occurred."
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

func handleErrors(c *gin.Context, trans ut.Translator) dto.ProblemDetail {
	for _, err := range c.Errors {
		if problemDetail, ok := handleMessageNotReadableError(err, c); ok {
			return problemDetail
		}
		if problemDetail, ok := handleValidationErrors(err, c, trans); ok {
			return problemDetail
		}
		if problemDetail, ok := handleConstraintViolationErrors(err, c); ok {
			return problemDetail
		}
		if problemDetail, ok := handleInvalidCredentialsErrors(err, c); ok {
			return problemDetail
		}
		if problemDetail, ok := handleJwtErrors(err, c); ok {
			return problemDetail
		}
		if problemDetail, ok := handleAccessDeniedErrors(err, c); ok {
			return problemDetail
		}
		if problemDetail, ok := handleConflictErrors(err, c); ok {
			return problemDetail
		}
		if problemDetail, ok := handleNotFoundErrors(err, c); ok { // Yeni hata tipi
			return problemDetail
		}
	}

	return handleInternalServerError(c)
}

func handleMessageNotReadableError(err *gin.Error, c *gin.Context) (dto.ProblemDetail, bool) {
	var messageNotReadableErr *customError.MessageNotReadableError
	if errors.As(err.Err, &messageNotReadableErr) {
		return dto.ProblemDetail{
			Type:     TypeAboutBlank,
			Title:    TitleBadRequest,
			Status:   http.StatusBadRequest,
			Detail:   "The request message could not be read. Please check the format and try again.",
			Error:    ErrorInvalidRequest,
			Instance: c.Request.URL.Path,
		}, true
	}
	return dto.ProblemDetail{}, false
}

func handleValidationErrors(err *gin.Error, c *gin.Context, trans ut.Translator) (dto.ProblemDetail, bool) {
	var validationErrs validator.ValidationErrors
	if errors.As(err.Err, &validationErrs) {
		var violations []dto.Violation
		for _, ve := range validationErrs {
			translatedMessage := ve.Translate(trans)
			violations = append(violations, dto.Violation{
				Code:          ve.Tag(),
				Object:        ve.StructNamespace(),
				Field:         ve.Field(),
				RejectedValue: ve.Param(),
				Message:       translatedMessage,
			})
		}
		return dto.ProblemDetail{
			Type:       TypeAboutBlank,
			Title:      TitleBadRequest,
			Status:     http.StatusBadRequest,
			Detail:     DetailValidationError,
			Error:      ErrorInvalidRequest,
			Instance:   c.Request.URL.Path,
			Violations: violations,
		}, true
	}
	return dto.ProblemDetail{}, false
}

func handleConstraintViolationErrors(err *gin.Error, c *gin.Context) (dto.ProblemDetail, bool) {
	var constraintErr customError.ConstraintViolationError
	if errors.As(err.Err, &constraintErr) {
		return dto.ProblemDetail{
			Type:       TypeAboutBlank,
			Title:      TitleBadRequest,
			Status:     http.StatusBadRequest,
			Detail:     DetailValidationError,
			Error:      ErrorInvalidRequest,
			Instance:   c.Request.URL.Path,
			Violations: constraintErr.Violations,
		}, true
	}
	return dto.ProblemDetail{}, false
}

func handleInvalidCredentialsErrors(err *gin.Error, c *gin.Context) (dto.ProblemDetail, bool) {
	var invalidCredentialsErr *customError.InvalidCredentialsError
	if errors.As(err.Err, &invalidCredentialsErr) {
		return dto.ProblemDetail{
			Type:     TypeAboutBlank,
			Title:    TitleUnauthorized,
			Status:   http.StatusUnauthorized,
			Detail:   "Invalid username or password.",
			Error:    "invalid_credentials",
			Instance: c.Request.URL.Path,
		}, true
	}
	return dto.ProblemDetail{}, false
}

func handleJwtErrors(err *gin.Error, c *gin.Context) (dto.ProblemDetail, bool) {
	var jwtErr *customError.JwtError
	if errors.As(err.Err, &jwtErr) {
		return dto.ProblemDetail{
			Type:     TypeAboutBlank,
			Title:    TitleUnauthorized,
			Status:   http.StatusUnauthorized,
			Detail:   "Invalid token.",
			Error:    "invalid_token",
			Instance: c.Request.URL.Path,
		}, true
	}
	return dto.ProblemDetail{}, false
}

func handleAccessDeniedErrors(err *gin.Error, c *gin.Context) (dto.ProblemDetail, bool) {
	var accessDeniedErr *customError.AccessDeniedError
	if errors.As(err.Err, &accessDeniedErr) {
		return dto.ProblemDetail{
			Type:     TypeAboutBlank,
			Title:    TitleAccessDenied,
			Status:   http.StatusForbidden,
			Detail:   "Access denied.",
			Error:    ErrorAccessDenied,
			Instance: c.Request.URL.Path,
		}, true
	}
	return dto.ProblemDetail{}, false
}

func handleConflictErrors(err *gin.Error, c *gin.Context) (dto.ProblemDetail, bool) {
	var conflictErr *customError.ResourceConflictError
	if errors.As(err.Err, &conflictErr) {
		return dto.ProblemDetail{
			Type:     TypeAboutBlank,
			Title:    TitleConflict,
			Status:   http.StatusConflict,
			Detail:   conflictErr.Error(),
			Error:    ErrorResourceConflict,
			Instance: c.Request.URL.Path,
		}, true
	}
	return dto.ProblemDetail{}, false
}

func handleNotFoundErrors(err *gin.Error, c *gin.Context) (dto.ProblemDetail, bool) {
	var notFoundErr *customError.ResourceNotFoundError
	if errors.As(err.Err, &notFoundErr) {
		return dto.ProblemDetail{
			Type:     TypeAboutBlank,
			Title:    TitleNotFound,
			Status:   http.StatusNotFound,
			Detail:   notFoundErr.Error(),
			Error:    ErrorResourceNotFound,
			Instance: c.Request.URL.Path,
		}, true
	}
	return dto.ProblemDetail{}, false
}

func handleInternalServerError(c *gin.Context) dto.ProblemDetail {
	return dto.ProblemDetail{
		Type:     TypeAboutBlank,
		Title:    TitleInternalServerError,
		Status:   http.StatusInternalServerError,
		Detail:   "An internal server error occurred. Please try again later.",
		Error:    ErrorInternalServer,
		Instance: c.Request.URL.Path,
	}
}
