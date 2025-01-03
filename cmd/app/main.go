package main

import (
	"gin-samples/internal/di"
)

func main() {
	container := di.NewContainer()
	container.Router.Run(":8080")
}
