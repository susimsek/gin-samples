package controller

import (
	"gin-samples/internal/dto"
	customError "gin-samples/internal/error"
	"gin-samples/internal/service"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HelloController interface {
	Hello(c *gin.Context)
	CreateGreeting(c *gin.Context)
	GetAllGreetings(c *gin.Context)
	GetGreetingByID(c *gin.Context)
	UpdateGreeting(c *gin.Context)
	DeleteGreeting(c *gin.Context)
}

type helloControllerImpl struct {
	HelloService service.HelloService
	Validator    *validator.Validate
	Trans        ut.Translator
}

func NewHelloController(service service.HelloService,
	validator *validator.Validate,
	trans ut.Translator) HelloController {
	return &helloControllerImpl{
		HelloService: service,
		Validator:    validator,
		Trans:        trans,
	}
}

// Hello godoc
// @Summary Get a greeting message
// @Description Returns a greeting message
// @Tags hello
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.GreetingResponse
// @Failure 401 {object} dto.ProblemDetail
// @Failure 500 {object} dto.ProblemDetail
// @Router /api/hello [get]
func (h *helloControllerImpl) Hello(c *gin.Context) {
	greeting := h.HelloService.GetGreeting()
	c.JSON(http.StatusOK, greeting)
}

// CreateGreeting godoc
// @Summary Create a new greeting message
// @Description Creates a new greeting
// @Tags hello
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body dto.GreetingInput true "Greeting Input"
// @Success 201 {object} dto.GreetingResponse
// @Failure 400 {object} dto.ProblemDetail
// @Failure 401 {object} dto.ProblemDetail
// @Failure 409 {object} dto.ProblemDetail
// @Failure 500 {object} dto.ProblemDetail
// @Router /api/hello [post]
func (h *helloControllerImpl) CreateGreeting(c *gin.Context) {
	var input dto.GreetingInput

	if err := c.ShouldBindJSON(&input); err != nil {
		_ = c.Error(&customError.MessageNotReadableError{
			Detail: err.Error(),
		})
		return
	}

	if err := h.Validator.Struct(input); err != nil {
		_ = c.Error(err)
		return
	}

	newGreeting, err := h.HelloService.CreateGreeting(input)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, newGreeting)
}

// GetAllGreetings godoc
// @Summary Get all greeting messages
// @Description Returns all greeting messages
// @Tags hello
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} dto.GreetingResponse
// @Failure 401 {object} dto.ProblemDetail
// @Failure 500 {object} dto.ProblemDetail
// @Router /api/hello/all [get]
func (h *helloControllerImpl) GetAllGreetings(c *gin.Context) {
	greetings, err := h.HelloService.GetAllGreetings()
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, greetings)
}

// GetGreetingByID godoc
// @Summary Get a greeting by ID
// @Description Returns a single greeting message by its ID
// @Tags hello
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Greeting ID"
// @Success 200 {object} dto.GreetingResponse
// @Failure 400 {object} dto.ProblemDetail
// @Failure 401 {object} dto.ProblemDetail
// @Failure 404 {object} dto.ProblemDetail
// @Failure 500 {object} dto.ProblemDetail
// @Router /api/hello/{id} [get]
func (h *helloControllerImpl) GetGreetingByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)

	// Handle invalid ID as a ConstraintViolationError
	if err != nil || id < 1 {
		violation := dto.Violation{
			Code:          "min",
			Field:         "id",
			RejectedValue: idParam,
			Message:       "ID must be a valid integer greater than or equal to 1",
		}
		constraintErr := customError.ConstraintViolationError{
			Violations: []dto.Violation{violation},
		}
		_ = c.Error(constraintErr)
		return
	}

	greeting, err := h.HelloService.GetGreetingByID(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, greeting)
}

// UpdateGreeting godoc
// @Summary Update a greeting message by ID
// @Description Updates a greeting message by its ID
// @Tags hello
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Greeting ID"
// @Param input body dto.GreetingInput true "Greeting Input"
// @Success 200 {object} dto.GreetingResponse
// @Failure 400 {object} dto.ProblemDetail
// @Failure 401 {object} dto.ProblemDetail
// @Failure 404 {object} dto.ProblemDetail
// @Failure 500 {object} dto.ProblemDetail
// @Router /api/hello/{id} [put]
func (h *helloControllerImpl) UpdateGreeting(c *gin.Context) {
	// Parse and validate ID from path
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil || id < 1 {
		violation := dto.Violation{
			Code:          "min",
			Field:         "id",
			RejectedValue: idParam,
			Message:       "ID must be a valid integer greater than or equal to 1",
		}
		constraintErr := customError.ConstraintViolationError{
			Violations: []dto.Violation{violation},
		}
		_ = c.Error(constraintErr)
		return
	}

	// Parse and validate input body
	var input dto.GreetingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		_ = c.Error(&customError.MessageNotReadableError{
			Detail: err.Error(),
		})
		return
	}

	if err := h.Validator.Struct(input); err != nil {
		_ = c.Error(err)
		return
	}

	// Call service to update the greeting
	updatedGreeting, err := h.HelloService.UpdateGreeting(uint(id), input)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, updatedGreeting)
}

// DeleteGreeting godoc
// @Summary Delete a greeting message by ID
// @Description Deletes a greeting message by its ID
// @Tags hello
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Greeting ID"
// @Success 204 "No Content"
// @Failure 400 {object} dto.ProblemDetail
// @Failure 401 {object} dto.ProblemDetail
// @Failure 404 {object} dto.ProblemDetail
// @Failure 500 {object} dto.ProblemDetail
// @Router /api/hello/{id} [delete]
func (h *helloControllerImpl) DeleteGreeting(c *gin.Context) {
	// Parse and validate ID from path
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil || id < 1 {
		violation := dto.Violation{
			Code:          "min",
			Field:         "id",
			RejectedValue: idParam,
			Message:       "ID must be a valid integer greater than or equal to 1",
		}
		constraintErr := customError.ConstraintViolationError{
			Violations: []dto.Violation{violation},
		}
		_ = c.Error(constraintErr)
		return
	}

	// Call service to delete the greeting
	err = h.HelloService.DeleteGreeting(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}

	// Return no content status
	c.Status(http.StatusNoContent)
}
