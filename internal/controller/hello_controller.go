package controller

import (
	"gin-samples/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloController interface {
	Hello(c *gin.Context)
}

type helloControllerImpl struct {
	HelloService service.HelloService
}

func NewHelloController(service service.HelloService) HelloController {
	return &helloControllerImpl{HelloService: service}
}

func (h *helloControllerImpl) Hello(c *gin.Context) {
	greeting := h.HelloService.GetGreeting()
	c.JSON(http.StatusOK, greeting)
}
