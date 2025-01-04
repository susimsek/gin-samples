package di

import (
	"gin-samples/config"
	"gin-samples/internal/controller"
	"gin-samples/internal/repository"
	"gin-samples/internal/router"
	"gin-samples/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Container struct {
	HelloRepository repository.HelloRepository
	HelloService    service.HelloService
	HelloController controller.HelloController
	Router          *gin.Engine
	Validator       *validator.Validate
	Translator      ut.Translator
}

func NewContainer() *Container {
	// Repository
	helloRepository := repository.NewHelloRepository()

	// Service
	helloService := service.NewHelloService(helloRepository)

	// Validator and Translator
	validate, translator := config.NewValidator()

	// Controller
	helloController := controller.NewHelloController(helloService, validate, translator)

	// Router
	r := router.SetupRouter(helloController, translator)

	return &Container{
		HelloRepository: helloRepository,
		HelloService:    helloService,
		HelloController: helloController,
		Router:          r,
		Validator:       validate,
		Translator:      translator,
	}
}
