package main

import (
	"gin-samples/internal/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
