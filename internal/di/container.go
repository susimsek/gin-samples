package di

import (
	"gin-samples/config"
	customCache "gin-samples/internal/cache"
	"gin-samples/internal/controller"
	"gin-samples/internal/mapper"
	"gin-samples/internal/repository"
	"gin-samples/internal/router"
	"gin-samples/internal/security"
	"gin-samples/internal/service"
	"gin-samples/internal/util"
	"github.com/dgraph-io/ristretto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Container struct {
	Config                *config.Config
	Cache                 *ristretto.Cache
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
	// Initialize CacheConfig
	cache := config.InitCache()

	cacheManager := customCache.NewCacheManager(cache)
	// Repository
	db := config.DatabaseConfig.InitDB()
	helloRepository := repository.NewHelloRepository(db, cacheManager)
	userRepository := repository.NewUserRepository(db, cacheManager)

	// Clock
	clock := &util.RealClock{} // Use RealClock for production

	// Mapper
	helloMapper := mapper.NewHelloMapper()

	// JWT KeyPair
	signKeyPair, encKeyPair := config.JweTokenConfig.InitJweKeyPair()

	// Token Generator
	tokenGenerator := security.NewTokenGenerator(
		signKeyPair, encKeyPair, cfg.TokenDuration, cfg.TokenIssuer)

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
		Cache:                 cache,
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
