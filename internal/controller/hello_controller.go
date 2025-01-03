package controller

import (
	"gin-samples/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloController struct {
	HelloService service.HelloService
}

func NewHelloController(service service.HelloService) *HelloController {
	return &HelloController{HelloService: service}
}

func (h *HelloController) Hello(c *gin.Context) {
	greeting := h.HelloService.GetGreeting()
	c.JSON(http.StatusOK, greeting)
}
