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
// @Success 200 {object} dto.GreetingResponse
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
// @Param input body dto.GreetingInput true "Greeting Input"
// @Success 201 {object} dto.GreetingResponse
// @Failure 400 {object} dto.ProblemDetail
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
// @Success 200 {array} dto.GreetingResponse
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
// @Param id path int true "Greeting ID"
// @Success 200 {object} dto.GreetingResponse
// @Failure 400 {object} dto.ProblemDetail
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
