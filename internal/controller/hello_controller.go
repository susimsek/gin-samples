package controller

import (
	"gin-samples/internal/model"
	"gin-samples/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloController interface {
	Hello(c *gin.Context)
	CreateGreeting(c *gin.Context)
}

type helloControllerImpl struct {
	HelloService service.HelloService
}

func NewHelloController(service service.HelloService) HelloController {
	return &helloControllerImpl{HelloService: service}
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
// @Router /api/hello [post]
func (h *helloControllerImpl) CreateGreeting(c *gin.Context) {
	var input model.GreetingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newGreeting := h.HelloService.CreateGreeting(input)
	c.JSON(http.StatusCreated, newGreeting)
}
