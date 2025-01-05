package main

import (
	"gin-samples/config"
	_ "gin-samples/docs"
	"gin-samples/internal/di"
	"log"
)

// @title Gin Samples API
// @version 1.0
// @description This is a sample server for Gin application.
// @termsOfService http://example.com/terms/

// @contact.name API Support
// @contact.url http://example.com/contact
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
func main() {
	cfg := config.LoadConfig()
	container := di.NewContainer()
	run(":"+cfg.ServerPort, container)
}

func run(addr string, container *di.Container) {
	if err := container.Router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
