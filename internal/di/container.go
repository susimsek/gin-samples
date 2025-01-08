package di

import (
	"gin-samples/config"
	"gin-samples/internal/controller"
	"gin-samples/internal/mapper"
	"gin-samples/internal/repository"
	"gin-samples/internal/router"
	"gin-samples/internal/security"
	"gin-samples/internal/service"
	"gin-samples/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Container struct {
	Config                *config.Config
	DB                    *gorm.DB
	HelloRepository       repository.HelloRepository
	UserRepository        repository.UserRepository
	HelloMapper           mapper.HelloMapper
	HelloService          service.HelloService
	AuthenticationService service.AuthenticationService
	TokenGenerator        security.TokenGenerator
	HelloController       controller.HelloController
	AuthController        controller.AuthenticationController
	HealthController      controller.HealthController
	Router                *gin.Engine
	Validator             *validator.Validate
	Translator            ut.Translator
	Clock                 util.Clock
}

func NewContainer(cfg *config.Config) *Container {
	// Repository
	db := config.DatabaseConfig.InitDB()
	helloRepository := repository.NewHelloRepository(db)
	userRepository := repository.NewUserRepository(db)

	// Clock
	clock := &util.RealClock{} // Use RealClock for production

	// Mapper
	helloMapper := mapper.NewHelloMapper()

	jwtKeyPair := config.InitJwtKeyPair()
	// Token Generator
	tokenGenerator := security.NewTokenGenerator(jwtKeyPair, cfg.TokenDuration)

	// Services
	helloService := service.NewHelloService(helloRepository, helloMapper, clock)
	authService := service.NewAuthenticationService(userRepository, tokenGenerator)

	// Validator and Translator
	validate, translator := config.NewValidator()

	// Controllers
	helloController := controller.NewHelloController(helloService, validate, translator)
	authController := controller.NewAuthenticationController(authService, validate, translator)
	healthController := controller.NewHealthController()

	// Router
	r := router.SetupRouter(helloController, healthController,
		authController, translator, tokenGenerator)

	return &Container{
		Config:                cfg,
		DB:                    db,
		HelloRepository:       helloRepository,
		UserRepository:        userRepository,
		HelloMapper:           helloMapper,
		HelloService:          helloService,
		AuthenticationService: authService,
		TokenGenerator:        tokenGenerator,
		HelloController:       helloController,
		AuthController:        authController,
		HealthController:      healthController,
		Router:                r,
		Validator:             validate,
		Translator:            translator,
		Clock:                 clock,
	}
}
