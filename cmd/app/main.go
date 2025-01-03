package main

import (
	_ "gin-samples/docs"
	"gin-samples/internal/di"
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
	container := di.NewContainer()
	run(":8080", container)
}

func run(addr string, container *di.Container) {
	container.Router.Run(addr)
}
