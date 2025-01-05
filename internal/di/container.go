package di

import (
	"gin-samples/config"
	"gin-samples/internal/controller"
	"gin-samples/internal/repository"
	"gin-samples/internal/router"
	"gin-samples/internal/service"
	"gin-samples/internal/utils"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Container struct {
	DB               *gorm.DB
	HelloRepository  repository.HelloRepository
	HelloService     service.HelloService
	HelloController  controller.HelloController
	HealthController controller.HealthController
	Router           *gin.Engine
	Validator        *validator.Validate
	Translator       ut.Translator
	Clock            utils.Clock
}

func NewContainer() *Container {
	// Repository
	db := config.InitDB()
	helloRepository := repository.NewHelloRepository(db)

	// Clock
	clock := &utils.RealClock{} // Use RealClock for production

	// Service
	helloService := service.NewHelloService(helloRepository, clock)

	// Validator and Translator
	validate, translator := config.NewValidator()

	// Controllers
	helloController := controller.NewHelloController(helloService, validate, translator)
	healthController := controller.NewHealthController()

	// Router
	r := router.SetupRouter(helloController, healthController, translator)

	return &Container{
		DB:               db,
		HelloRepository:  helloRepository,
		HelloService:     helloService,
		HelloController:  helloController,
		HealthController: healthController,
		Router:           r,
		Validator:        validate,
		Translator:       translator,
		Clock:            clock,
	}
}
