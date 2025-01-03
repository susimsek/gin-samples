package main

import (
	"gin-samples/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
