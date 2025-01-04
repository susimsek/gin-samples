package controller

import (
	customError "gin-samples/internal/error"
	"gin-samples/internal/model"
	"gin-samples/internal/service"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloController interface {
	Hello(c *gin.Context)
	CreateGreeting(c *gin.Context)
	GetAllGreetings(c *gin.Context)
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
// @Success 200 {object} model.Greeting
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
// @Param input body model.GreetingInput true "Greeting Input"
// @Success 201 {object} model.Greeting
// @Failure 400 {object} model.ProblemDetail
// @Failure 409 {object} model.ProblemDetail
// @Failure 500 {object} model.ProblemDetail
// @Router /api/hello [post]
func (h *helloControllerImpl) CreateGreeting(c *gin.Context) {
	var input model.GreetingInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(&customError.MessageNotReadableError{
			Detail: err.Error(),
		})
		return
	}

	if err := h.Validator.Struct(input); err != nil {
		c.Error(err)
		return
	}

	// Call service layer
	newGreeting, err := h.HelloService.CreateGreeting(input)
	if err != nil {
		c.Error(err) // ErrorHandlingMiddleware will handle this
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
// @Success 200 {array} model.Greeting
// @Router /api/hello/all [get]
func (h *helloControllerImpl) GetAllGreetings(c *gin.Context) {
	greetings := h.HelloService.GetAllGreetings()
	c.JSON(http.StatusOK, greetings)
}
