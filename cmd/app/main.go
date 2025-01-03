package main

import (
	"gin-samples/internal/di"
)

func main() {
	container := di.NewContainer()
	run(":8080", container)
}

func run(addr string, container *di.Container) {
	container.Router.Run(addr)
}
