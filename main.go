// main.go
package main

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
	return r
}

func main() {
	router := SetupRouter()
	router.Run(":8080")
}
